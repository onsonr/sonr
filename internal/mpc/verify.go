package mpc

import (
	"errors"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// - The R and S values must be in the valid range for secp256k1 scalars:
//   - Negative values are rejected
//   - Zero is rejected
//   - Values greater than or equal to the secp256k1 group order are rejected
func DeserializeECDSASecp256k1Signature(sigStr []byte) (*crypto.MPCECDSASignature, error) {
	rBytes := sigStr[:33]
	sBytes := sigStr[33:65]

	sig := crypto.NewEmptyECDSASecp256k1Signature()
	if err := sig.R.UnmarshalBinary(rBytes); err != nil {
		return nil, errors.New("malformed signature: R is not in the range [1, N-1]")
	}

	// S must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	if err := sig.S.UnmarshalBinary(sBytes); err != nil {
		return nil, errors.New("malformed signature: S is not in the range [1, N-1]")
	}

	// Create and return the signature.
	return &sig, nil
}

// VerifyCMP verifies a message with the given public key using the CMP protocol.
func VerifyCMP(config *cmp.Config, m []byte, sig []byte) (bool, error) {
	sigObj, err := DeserializeECDSASecp256k1Signature(sig)
	if err != nil {
		return false, err
	}
	return sigObj.Verify(config.PublicPoint(), m), nil
}
