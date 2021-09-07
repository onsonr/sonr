package transfer

import (
	"io"

	s2 "github.com/klauspost/compress/s2"
	msg "github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/emitter"
)

type ItemWriter interface {
	Progress() []byte
	WriteTo(writer msg.WriteCloser) error
}

type itemWriter struct {
	ItemWriter
	emitter *emitter.Emitter
	index   int
	size    int
	item    *common.Transfer_Item
}

func newWriter(i *common.Transfer_Item, em *emitter.Emitter) ItemWriter {
	return &itemWriter{
		item:    i,
		emitter: em,
	}
}

func EncodeStream(src []byte, dst io.Writer) error {
	enc := s2.NewWriter(dst)
	// The encoder owns the buffer until Flush or Close is called.
	err := enc.EncodeBuffer(src)
	if err != nil {
		enc.Close()
		return err
	}
	// Blocks until compression is done.
	return enc.Close()
}
