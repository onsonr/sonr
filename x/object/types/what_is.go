package types

import (
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

func NewWhatIsFromBuf(cd *ot.WhatIs) *WhatIs {
	return &WhatIs{
		Did:       cd.GetDid(),
		Creator:   cd.GetCreator(),
		Timestamp: cd.GetTimestamp(),
		IsActive:  cd.GetIsActive(),
		ObjectDoc: NewObjectDocFromBuf(cd.GetObjectDoc()),
	}
}

func NewWhatIsToBuf(cd *WhatIs) *ot.WhatIs {
	return &ot.WhatIs{
		Did:       cd.GetDid(),
		Creator:   cd.GetCreator(),
		Timestamp: cd.GetTimestamp(),
		IsActive:  cd.GetIsActive(),
		ObjectDoc: NewObjectDocToBuf(cd.GetObjectDoc()),
	}
}
