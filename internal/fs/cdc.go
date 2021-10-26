package fs

import (
	"bytes"
	"errors"
	"io"
	"math"
	"os"
)

const (
	kiB = 1024
	miB = 1024 * kiB

	minSize = 64
	maxSize = 1 << 30

	defaultNormalization = 2
	interval             = 25
)

// Chunker implements the FastCDC content defined chunking algorithm.
// See https://www.usenix.org/system/files/conference/atc16/atc16-paper-xia.pdf.
type Chunker struct {
	minSize  int
	maxSize  int
	normSize int

	maskS uint64
	maskL uint64

	rd io.Reader

	buf    []byte
	cursor int
	offset int
	eof    bool
}

// ChunkerOptions configures the options for the Chunker.
type ChunkerOptions struct {
	// NormalSize is the target chunk size. Typically a power of 2. It must be in the
	// range 64B to 1GiB.
	AverageSize int

	// (Optional) MinSize is the minimum allowed chunk size. By default, it's set to
	// AverageSize / 4.
	MinSize int

	// (Optional) MaxSize is the maximum allowed chunk size. By default, it's set to
	// AverageSize * 4.
	MaxSize int

	// (Optional) Sets the chunk normalization level. It may be set to 1, 2 or 3,
	// unless DisableNormalization is set, in which case it's ignored. By default,
	// it's set to 2.
	Normalization int

	// (Optional) DisableNormalization turns normalization off. By default, it's set to
	// false.
	DisableNormalization bool

	// (Optional) Seed alters the lookup table of the rolling hash algorithm to mitigate
	// chunk-size based fingerprinting attacks. It may be set to a random uint64.
	Seed uint64

	// (Optional) BufSize is the size of the internal buffer used while chunking. It has
	// no effect on the chuking output, but performance is improved with larger buffers.
	// It must be at least MaxSize. Recommended values are 1 to 3 times MaxSize. By
	// default it is set to MaxSize * 2.
	BufSize int
}

func (opts *ChunkerOptions) setDefaults() {
	if opts.MinSize == 0 {
		opts.MinSize = opts.AverageSize / 4
	}
	if opts.MaxSize == 0 {
		opts.MaxSize = opts.AverageSize * 4
	}
	if opts.BufSize == 0 {
		opts.BufSize = opts.MaxSize * 2
	}
	if !opts.DisableNormalization && opts.Normalization == 0 {
		opts.Normalization = 2
	}
}

// Chunk stores a content-defined chunk returned by a Chunker.
type Chunk struct {
	// Offset is the number of bytes from the start of the reader to the beginning of
	// the chunk.
	Offset int

	// Length is the length of the chunk in bytes. Same as len(Data).
	Length int

	// Data is the chunk data.
	Data []byte

	// Fingerprint is the value of the rolling hash algorithm for the chunk data.
	Fingerprint uint64
}

// NewChunker returns a Chunker with the given Options.
func NewChunker(rd io.Reader, opts ChunkerOptions) (*Chunker, error) {
	opts.setDefaults()
	if err := opts.validate(); err != nil {
		return nil, err
	}

	for i := 0; i < len(table); i++ {
		table[i] = table[i] ^ opts.Seed
	}

	normalization := opts.Normalization
	if opts.DisableNormalization {
		normalization = 0
	}
	bits := int(math.Round(math.Log2(float64(opts.AverageSize))))
	smallBits := bits + normalization
	largeBits := bits - normalization

	chunker := &Chunker{
		minSize:  opts.MinSize,
		maxSize:  opts.MaxSize,
		normSize: opts.AverageSize,
		maskS:    (1 << smallBits) - 1,
		maskL:    (1 << largeBits) - 1,
		rd:       rd,
		buf:      make([]byte, opts.BufSize),
		cursor:   opts.BufSize,
	}
	return chunker, nil
}

func NewFileChunker(path string, size int64) (*Chunker, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var averageSize int
	if size < interval {
		averageSize = int(size)
	} else {
		averageSize = int(size / interval)
	}

	return NewChunker(bytes.NewReader(buf), ChunkerOptions{
		AverageSize: averageSize,
	})
}

func (c *Chunker) fillBuffer() error {
	n := len(c.buf) - c.cursor
	if n >= c.maxSize {
		return nil
	}

	// Move all data after the cursor to the start of the buffer
	copy(c.buf[:n], c.buf[c.cursor:])
	c.cursor = 0

	if c.eof {
		c.buf = c.buf[:n]
		return nil
	}

	// Fill the rest of the buffer
	m, err := io.ReadFull(c.rd, c.buf[n:])
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		c.buf = c.buf[:n+m]
		c.eof = true
	} else if err != nil {
		return err
	}
	return nil
}

// Next returns the next Chunk from the reader or io.EOF after the last chunk has been
// read. The chunk data is invalidated when Next is called again.
func (c *Chunker) Next() (Chunk, error) {
	if err := c.fillBuffer(); err != nil {
		return Chunk{}, err
	}
	if len(c.buf) == 0 {
		return Chunk{}, io.EOF
	}

	length, fp := c.nextChunk(c.buf[c.cursor:])
	chunk := Chunk{
		Offset:      c.offset,
		Length:      length,
		Data:        c.buf[c.cursor : c.cursor+length],
		Fingerprint: fp,
	}

	c.cursor += length
	c.offset += chunk.Length
	return chunk, nil
}

func (c *Chunker) nextChunk(data []byte) (int, uint64) {
	fp := uint64(0)
	i := c.minSize
	if len(data) <= c.minSize {
		return len(data), fp
	}

	n := min(len(data), c.maxSize)
	for ; i < min(n, c.normSize); i++ {
		fp = (fp << 1) + table[data[i]]
		if (fp & c.maskS) == 0 {
			return i + 1, fp
		}
	}

	for ; i < n; i++ {
		fp = (fp << 1) + table[data[i]]
		if (fp & c.maskL) == 0 {
			return i + 1, fp
		}
	}
	return i, fp
}

func (opts ChunkerOptions) validate() error {
	if opts.AverageSize == 0 {
		return errors.New("option AverageSize is required")
	}
	if opts.MinSize < minSize || opts.MinSize > maxSize {
		return errors.New("option MinSize must be in range 64B to 1GiB")
	}
	if opts.MaxSize < minSize || opts.MaxSize > maxSize {
		return errors.New("option MaxSize must be in range 64B to 1GiB")
	}
	if opts.MaxSize <= opts.MinSize {
		return errors.New("option MinSize must be less than option MaxSize")
	}
	if opts.AverageSize > opts.MaxSize || opts.AverageSize < opts.MinSize {
		return errors.New("option AverageSize must be betweeen MinSize and MaxSize")
	}
	if !opts.DisableNormalization && (opts.Normalization <= 0 || opts.Normalization > 4) {
		return errors.New("option Normalization must be 0, 1, 2 or 3")
	}
	if opts.BufSize <= opts.MaxSize {
		return errors.New("option BufSize, if specified, must be at least MaxSize")
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 256 random uint64s for the rolling hash function
var table [256]uint64 = [256]uint64{
	0xe80e8d55032474b3, 0x11b25b61f5924e15, 0x03aa5bd82a9eb669, 0xc45a153ef107a38c,
	0xeac874b86f0f57b9, 0xa5ccedec95ec79c7, 0xe15a3320ad42ac0a, 0x5ed3583fa63cec15,
	0xcd497bf624a4451d, 0xf9ade5b059683605, 0x773940c03fb11ca1, 0xa36b16e4a6ae15b2,
	0x67afd1adb5a89eac, 0xc44c75ee32f0038e, 0x2101790f365c0967, 0x76415c64a222fc4a,
	0x579929249a1e577a, 0xe4762fc41fdbf750, 0xea52198e57dfcdcc, 0xe2535aafe30b4281,
	0xcb1a1bd6c77c9056, 0x5a1aa9bfc4612a62, 0x15a728aef8943eb5, 0x2f8f09738a8ec8d9,
	0x200f3dec9fac8074, 0x0fa9a7b1e0d318df, 0x06c0804ffd0d8e3a, 0x630cbc412669dd25,
	0x10e34f85f4b10285, 0x2a6fe8164b9b6410, 0xcacb57d857d55810, 0x77f8a3a36ff11b46,
	0x66af517e0dc3003e, 0x76c073c789b4009a, 0x853230dbb529f22a, 0x1e9e9c09a1f77e56,
	0x1e871223802ee65d, 0x37fe4588718ff813, 0x10088539f30db464, 0x366f7470b80b72d1,
	0x33f2634d9a6b31db, 0xd43917751d69ea18, 0xa0f492bc1aa7b8de, 0x3f94e5a8054edd20,
	0xedfd6e25eb8b1dbf, 0x759517a54f196a56, 0xe81d5006ec7b6b17, 0x8dd8385fa894a6b7,
	0x45f4d5467b0d6f91, 0xa1f894699de22bc8, 0x33829d09ef93e0fe, 0x3e29e250caed603c,
	0xf7382cba7f63a45e, 0x970f95412bb569d1, 0xc7fcea456d356b4b, 0x723042513f3e7a57,
	0x17ae7688de3596f1, 0x27ac1fcd7cd23c1a, 0xf429beeb78b3f71f, 0xd0780692fb93a3f9,
	0x9f507e28a7c9842f, 0x56001ad536e433ae, 0x7e1dd1ecf58be306, 0x15fee353aa233fc6,
	0xb033a0730b7638e8, 0xeb593ad6bd2406d1, 0x7c86502574d0f133, 0xce3b008d4ccb4be7,
	0xf8566e3d383594c8, 0xb2c261e9b7af4429, 0xf685e7e253799dbb, 0x05d33ed60a494cbc,
	0xeaf88d55a4cb0d1a, 0x3ee9368a902415a1, 0x8980fe6a8493a9a4, 0x358ed008cb448631,
	0xd0cb7e37b46824b8, 0xe9bc375c0bc94f84, 0xea0bf1d8e6b55bb3, 0xb66a60d0f9f6f297,
	0x66db2cc4807b3758, 0x7e4e014afbca8b4d, 0xa5686a4938b0c730, 0xa5f0d7353d623316,
	0x26e38c349242d5e8, 0xeeefa80a29858e30, 0x8915cb912aa67386, 0x4b957a47bfc420d4,
	0xbb53d051a895f7e1, 0x09f5e3235f6911ce, 0x416b98e695cfb7ce, 0x97a08183344c5c86,
	0xbf68e0791839a861, 0xea05dde59ed3ed56, 0x0ca732280beda160, 0xac748ed62fe7f4e2,
	0xc686da075cf6e151, 0xe1ba5658f4af05c8, 0xe9ff09fbeb67cc35, 0xafaea9470323b28d,
	0x0291e8db5bb0ac2a, 0x342072a9bbee77ae, 0x03147eed6b3d0a9c, 0x21379d4de31dbadb,
	0x2388d965226fb986, 0x52c96988bfebabfa, 0xa6fc29896595bc2d, 0x38fa4af70aa46b8b,
	0xa688dd13939421ee, 0x99d5275d9b1415da, 0x453d31bb4fe73631, 0xde51debc1fbe3356,
	0x75a3c847a06c622f, 0xe80e32755d272579, 0x5444052250d8ec0d, 0x8f17dfda19580a3b,
	0xf6b3e9363a185e42, 0x7a42adec6868732f, 0x32cb6a07629203a2, 0x1eca8957defe56d9,
	0x9fa85e4bc78ff9ed, 0x20ff07224a499ca7, 0x3fa6295ff9682c70, 0xe3d5b1e3ce993eff,
	0xa341209362e0b79a, 0x64bd9eae5712ffe8, 0xceebb537babbd12a, 0x5586ef404315954f,
	0x46c3085c938ab51a, 0xa82ccb9199907cee, 0x8c51b6690a3523c8, 0xc4dbd4c9ae518332,
	0x979898dbb23db7b2, 0x1b5b585e6f672a9d, 0xce284da7c4903810, 0x841166e8bb5f1c4f,
	0xb7d884a3fceca7d0, 0xa76468f5a4572374, 0xc10c45f49ee9513d, 0x68f9a5663c1908c9,
	0x0095a13476a6339d, 0xd1d7516ffbe9c679, 0xfd94ab0c9726f938, 0x627468bbdb27c959,
	0xedc3f8988e4a8c9a, 0x58efd33f0dfaa499, 0x21e37d7e2ef4ac8b, 0x297f9ab5586259c6,
	0xda3ba4dc6cb9617d, 0xae11d8d9de2284d2, 0xcfeed88cb3729865, 0xefc2f9e4f03e2633,
	0x8226393e8f0855a4, 0xd6e25fd7acf3a767, 0x435784c3bfd6d14a, 0xf97142e6343fe757,
	0xd73b9fe826352f85, 0x6c3ac444b5b2bd76, 0xd8e88f3e9fd4a3fd, 0x31e50875c36f3460,
	0xa824f1bf88cf4d44, 0x54a4d2c8f5f25899, 0xbff254637ce3b1e6, 0xa02cfe92561b3caa,
	0x7bedb4edee9f0af7, 0x879c0620ac49a102, 0xa12c4ccd23b332e7, 0x09a5ff47bf94ed1e,
	0x7b62f43cd3046fa0, 0xaa3af0476b9c2fb9, 0x22e55301abebba8e, 0x3a6035c42747bd58,
	0x1705373106c8ec07, 0xb1f660de828d0628, 0x065fe82d89ca563d, 0xf555c2d8074d516d,
	0x6bb6c186b423ee99, 0x54a807be6f3120a8, 0x8a3c7fe2f88860b8, 0xbeffc344f5118e81,
	0xd686e80b7d1bd268, 0x661aef4ef5e5e88b, 0x5bf256c654cd1dda, 0x9adb1ab85d7640f4,
	0x68449238920833a2, 0x843279f4cebcb044, 0xc8710cdefa93f7bb, 0x236943294538f3e6,
	0x80d7d136c486d0b4, 0x61653956b28851d3, 0x3f843be9a9a956b5, 0xf73cfbbf137987e5,
	0xcf0cb6dee8ceac2c, 0x50c401f52f185cae, 0xbdbe89ce735c4c1c, 0xeef3ade9c0570bc7,
	0xbe8b066f8f64cbf6, 0x5238d6131705dcb9, 0x20219086c950e9f6, 0x634468d9ed74de02,
	0x0aba4b3d705c7fa5, 0x3374416f725a6672, 0xe7378bdf7beb3bc6, 0x0f7b6a1b1cee565b,
	0x234e4c41b0c33e64, 0x4efa9a0c3f21fe28, 0x1167fc551643e514, 0x9f81a69d3eb01fa4,
	0xdb75c22b12306ed0, 0xe25055d738fc9686, 0x9f9f167a3f8507bb, 0x195f8336d3fbe4d3,
	0x8442b6feffdcb6f6, 0x1e07ed24746ffde9, 0x140e31462d555266, 0x8bd0ce515ae1406e,
	0x2c0be0042b5584b3, 0x35a23d0e15d45a60, 0xc14f1ba147d9bc83, 0xbbf168691264b23f,
	0xad2cc7b57e589ade, 0x9501963154c7815c, 0x9664afa6b8d67d47, 0x7f9e5101fea0a81c,
	0x45ecffb610d25bfd, 0x3157f7aecf9b6ab3, 0xc43ca6f88d87501d, 0x9576ff838dee38dc,
	0x93f21afe0ce1c7d7, 0xceac699df343d8f9, 0x2fec49e29f03398d, 0x8805ccd5730281ed,
	0xf9fc16fc750a8e59, 0x35308cc771adf736, 0x4a57b7c9ee2b7def, 0x03a4c6cdc937a02a,
	0x6c9a8a269fc8c4fc, 0x4681decec7a03f43, 0x342eecded1353ef9, 0x8be0552d8413a867,
	0xc7b4ac51beda8be8, 0xebcc64fb719842c0, 0xde8e4c7fb6d40c1c, 0xcc8263b62f9738b1,
	0xd3cfc0f86511929a, 0x466024ce8bb226ea, 0x459ff690253a3c18, 0x98b27e9d91284c9c,
	0x75c3ae8aa3af373d, 0xfbf8f8e79a866ffc, 0x32327f59d0662799, 0x8228b57e729e9830,
	0x065ceb7a18381b58, 0xd2177671a31dc5ff, 0x90cd801f2f8701f9, 0x9d714428471c65fe,
}
