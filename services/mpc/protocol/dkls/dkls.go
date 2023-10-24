package dkls

import (
	"fmt"

	mpcv1types "github.com/sonr-io/core/services/mpc/types"
	"github.com/sonr-io/kryptology/pkg/core/curves"
	"github.com/sonr-io/kryptology/pkg/core/protocol"
	dklsv1 "github.com/sonr-io/kryptology/pkg/tecdsa/dkls/v1"
)

// The DKLSKeygen function generates a set of keyshares.
func DKLSKeygen() (mpcv1types.KeyshareSet, error) {
	curve := curves.K256()
	aliceDkg := dklsv1.NewAliceDkg(curve, protocol.Version1)
	bobDkg := dklsv1.NewBobDkg(curve, protocol.Version1)
	aErr, bErr := mpcv1types.RunIteratedProtocol(bobDkg, aliceDkg)
	if aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error running protocol: aErr=%v, bErr=%v", aErr, bErr)
	}
	aliceDkgResultMessage, err := aliceDkg.Result(protocol.Version1)
	if err != nil {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error getting Alice DKG result: %v", err)
	}
	bobDkgResultMessage, err := bobDkg.Result(protocol.Version1)
	if err != nil {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error getting Bob DKG result: %v", err)
	}
	return mpcv1types.NewKeyshareSet(aliceDkgResultMessage, bobDkgResultMessage), nil
}

// The DKLSRefresh function performs a key refresh protocol between two participants, Alice and Bob, and returns the updated keyshare set.
func DKLSRefresh(ksSet mpcv1types.KeyshareSet) (mpcv1types.KeyshareSet, error) {
	curve := curves.K256()
	aliceRefresh, err := dklsv1.NewAliceRefresh(curve, ksSet.DKGAtIndex(0), protocol.Version1)
	if err != nil {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error creating Alice refresh: %v", err)
	}
	bobRefresh, err := dklsv1.NewBobRefresh(curve, ksSet.DKGAtIndex(1), protocol.Version1)
	if err != nil {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error creating Bob refresh: %v", err)
	}
	aErr, bErr := mpcv1types.RunIteratedProtocol(aliceRefresh, bobRefresh)
	if aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error running protocol: aErr=%v, bErr=%v", aErr, bErr)
	}
	aliceRefreshResultMessage, err := aliceRefresh.Result(protocol.Version1)
	if err != nil {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error getting Alice refresh result: %v", err)
	}
	bobRefreshResultMessage, err := bobRefresh.Result(protocol.Version1)
	if err != nil {
		return mpcv1types.EmptyKeyshareSet(), fmt.Errorf("error getting Bob refresh result: %v", err)
	}
	return mpcv1types.NewKeyshareSet(aliceRefreshResultMessage, bobRefreshResultMessage), nil
}
