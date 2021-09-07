package transfer

import (
	"io"

	s2 "github.com/klauspost/compress/s2"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/emitter"
)

type ItemReader interface {
	Progress() []byte
	ReadFrom(src io.Reader, dst io.Writer) (int64, error)
}

type itemReader struct {
	ItemReader
	emitter *emitter.Emitter
	index   int
	size    int
	total   int
	item    *common.Transfer_Item
}

func newReader(i *common.Transfer_Item, em *emitter.Emitter) ItemReader {
	return &itemReader{
		item:    i,
		emitter: em,
	}
}

func (ir *itemReader) ReadFrom(src io.Reader, dst io.Writer) (int64, error) {
	dec := s2.NewReader(src)
	n, err := io.Copy(dst, dec)
	if err != nil {
		return n, err
	}
	ir.emitter.Emit(emitter.EMIT_PROGRESS_EVENT, n)
	return n, nil
}
