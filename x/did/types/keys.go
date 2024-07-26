package types

import (
	fmt "fmt"
	"math/big"

	"cosmossdk.io/collections"
	"github.com/onsonr/crypto/core/curves"
	"github.com/onsonr/crypto/signatures/ecdsa"
	"golang.org/x/crypto/sha3"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
)

// ParamsKey saves the current module params.
var ParamsKey = collections.NewPrefix(0)

const (
	ModuleName = "did"

	StoreKey = ModuleName

	QuerierRoute = ModuleName
)

var ORMModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: "did/v1/state.proto"},
	},
	Prefix: []byte{0},
}

// VerifySignature verifies the signature of a message
func VerifySignature(key []byte, msg []byte, sig []byte) bool {
	pp, err := BuildEcPoint(key)
	if err != nil {
		return false
	}
	sigEd, err := ecdsa.DeserializeSecp256k1Signature(sig)
	if err != nil {
		return false
	}
	hash := sha3.New256()
	_, err = hash.Write(msg)
	if err != nil {
		return false
	}
	digest := hash.Sum(nil)
	return curves.VerifyEcdsa(pp, digest[:], sigEd)
}

// BuildEcPoint builds an elliptic curve point from a compressed byte slice
func BuildEcPoint(pubKey []byte) (*curves.EcPoint, error) {
	crv := curves.K256()
	x := new(big.Int).SetBytes(pubKey[1:33])
	y := new(big.Int).SetBytes(pubKey[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}
