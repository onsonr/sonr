package types

import (
	ct "go.buf.build/grpc/go/sonr-io/blockchain/channel"
)

func NewHowIsFromBuf(cd *ct.HowIs) *HowIs {
	return &HowIs{
		Did:       cd.GetDid(),
		Creator:   cd.GetCreator(),
		Timestamp: cd.GetTimestamp(),
		IsActive:  cd.GetIsActive(),
		Channel:   NewChannelDocFromBuf(cd.GetChannel()),
	}
}

func NewHowIsToBuf(cd *HowIs) *ct.HowIs {
	return &ct.HowIs{
		Did:       cd.GetDid(),
		Creator:   cd.GetCreator(),
		Timestamp: cd.GetTimestamp(),
		IsActive:  cd.GetIsActive(),
		Channel:   NewChannelDocToBuf(cd.GetChannel()),
	}
}
