package transfer

import (
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
)

const maxFileBufferSize = 64

type TransferQueue struct {
	// Channels
	incomingCh chan *IncomingFile
	outgoingCh chan *sf.ProcessedFile

	// Processed
	processed []*sf.ProcessedFile
}

func StartQueue() *TransferQueue {
	return &TransferQueue{
		incomingCh: make(chan *IncomingFile, maxFileBufferSize),
		outgoingCh: make(chan *sf.ProcessedFile, maxFileBufferSize),
		processed:  make([]*sf.ProcessedFile, maxFileBufferSize),
	}
}

func (tq *TransferQueue) Processed(pf *sf.ProcessedFile) {
	tq.processed = append(tq.processed, pf)
}

func (tq *TransferQueue) NewIncoming(inv *md.AuthInvite, dirs *md.Directories, call md.TransferCallback) {
	incFile := NewIncomingFile(inv, dirs, call)
	tq.incomingCh <- incFile
}

func (tq *TransferQueue) NewOutgoing(inv *md.AuthInvite, dirs *md.Directories, call md.TransferCallback) {
	incFile := NewIncomingFile(inv, dirs, call)
	tq.incomingCh <- incFile
}
