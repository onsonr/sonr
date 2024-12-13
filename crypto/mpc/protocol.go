package mpc

import (
	"github.com/ipfs/kubo/client/rpc"
	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"github.com/onsonr/sonr/crypto/tecdsa/dklsv1"
)

// GenIPFSEnclave generates a new MPC keyshare
func GenIPFSEnclave(ipc *rpc.HttpApi) (KeyEnclave, error) {
	curve := curves.K256()
	valKs := dklsv1.NewAliceDkg(curve, protocol.Version1)
	userKs := dklsv1.NewBobDkg(curve, protocol.Version1)
	aErr, bErr := RunProtocol(userKs, valKs)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	valRes, err := valKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	valShare, err := EncodeKeyshare(valRes, RoleValidator)
	if err != nil {
		return nil, err
	}
	userRes, err := userKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userShare, err := EncodeKeyshare(userRes, RoleUser)
	if err != nil {
		return nil, err
	}
	return initKeyEnclave(valShare, userShare)
}

// GenEnclave generates a new MPC keyshare
func GenEnclave() (KeyEnclave, error) {
	curve := curves.K256()
	valKs := dklsv1.NewAliceDkg(curve, protocol.Version1)
	userKs := dklsv1.NewBobDkg(curve, protocol.Version1)
	aErr, bErr := RunProtocol(userKs, valKs)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	valRes, err := valKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	valShare, err := EncodeKeyshare(valRes, RoleValidator)
	if err != nil {
		return nil, err
	}
	userRes, err := userKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userShare, err := EncodeKeyshare(userRes, RoleUser)
	if err != nil {
		return nil, err
	}
	return initKeyEnclave(valShare, userShare)
}

// ExecuteSigning runs the MPC signing protocol
func ExecuteSigning(signFuncVal SignFunc, signFuncUser SignFunc) ([]byte, error) {
	aErr, bErr := RunProtocol(signFuncVal, signFuncUser)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	out, err := signFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	s, err := dklsv1.DecodeSignature(out)
	if err != nil {
		return nil, err
	}
	sig, err := serializeSignature(s)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// ExecuteRefresh runs the MPC refresh protocol
func ExecuteRefresh(refreshFuncVal RefreshFunc, refreshFuncUser RefreshFunc) (KeyEnclave, error) {
	aErr, bErr := RunProtocol(refreshFuncVal, refreshFuncUser)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	valRefreshResult, err := refreshFuncVal.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	valShare, err := EncodeKeyshare(valRefreshResult, RoleValidator)
	if err != nil {
		return nil, err
	}
	userRefreshResult, err := refreshFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userShare, err := EncodeKeyshare(userRefreshResult, RoleUser)
	if err != nil {
		return nil, err
	}
	return initKeyEnclave(valShare, userShare)
}

// For DKG bob starts first. For refresh and sign, Alice starts first.
func RunProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) (error, error) {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return nil, bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return aErr, nil
		}
	}
	return aErr, bErr
}
