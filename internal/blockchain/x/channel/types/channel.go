package types

import (
	ot "github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	ct "go.buf.build/grpc/go/sonr-io/blockchain/channel"
)

func NewChannelDocFromBuf(cd *ct.ChannelDoc) *ChannelDoc {
	return &ChannelDoc{
		Did:              cd.GetDid(),
		Label:            cd.GetLabel(),
		Description:      cd.GetDescription(),
		RegisteredObject: ot.NewObjectDocFromBuf(cd.GetRegisteredObject()),
	}
}

func NewChannelDocToBuf(cd *ChannelDoc) *ct.ChannelDoc {
	return &ct.ChannelDoc{
		Did:              cd.GetDid(),
		Label:            cd.GetLabel(),
		Description:      cd.GetDescription(),
		RegisteredObject: ot.NewObjectDocToBuf(cd.GetRegisteredObject()),
	}
}
