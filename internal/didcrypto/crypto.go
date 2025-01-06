package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

// GenerateEd25519 generates an Ed25519 private key and the matching DID.
// This is the RECOMMENDED algorithm.
func GenerateEd25519() (crypto.PrivKey, DID, error) {
	priv, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, Undef, nil
	}
	did, err := FromPubKey(pub)
	return priv, did, err
}

// GenerateRSA generates a RSA private key and the matching DID.
func GenerateRSA() (crypto.PrivKey, DID, error) {
	// NIST Special Publication 800-57 Part 1 Revision 5
	// Section 5.6.1.1 (Table 2)
	// Paraphrased: 2048-bit RSA keys are secure until 2030 and 3072-bit keys are recommended for longer-term security.
	const keyLength = 3072

	priv, pub, err := crypto.GenerateRSAKeyPair(keyLength, rand.Reader)
	if err != nil {
		return nil, Undef, nil
	}
	did, err := FromPubKey(pub)
	return priv, did, err
}

// GenerateSecp256k1 generates a Secp256k1 private key and the matching DID.
func GenerateSecp256k1() (crypto.PrivKey, DID, error) {
	priv, pub, err := crypto.GenerateSecp256k1Key(rand.Reader)
	if err != nil {
		return nil, Undef, nil
	}
	did, err := FromPubKey(pub)
	return priv, did, err
}

// GenerateECDSA generates an ECDSA private key and the matching DID
// for the default P256 curve.
func GenerateECDSA() (crypto.PrivKey, DID, error) {
	return GenerateECDSAWithCurve(P256)
}

// GenerateECDSAWithCurve generates an ECDSA private key and matching
// DID for the user-supplied curve
func GenerateECDSAWithCurve(code multicodec.Code) (crypto.PrivKey, DID, error) {
	var curve elliptic.Curve

	switch code {
	case P256:
		curve = elliptic.P256()
	case P384:
		curve = elliptic.P384()
	case P521:
		curve = elliptic.P521()
	default:
		return nil, Undef, errors.New("unsupported ECDSA curve")
	}

	priv, pub, err := crypto.GenerateECDSAKeyPairWithCurve(curve, rand.Reader)
	if err != nil {
		return nil, Undef, err
	}

	did, err := FromPubKey(pub)

	return priv, did, err
}

// FromPrivKey is a convenience function that returns the DID associated
// with the public key associated with the provided private key.
func FromPrivKey(privKey crypto.PrivKey) (DID, error) {
	return FromPubKey(privKey.GetPublic())
}

// FromPubKey returns a did:key constructed from the provided public key.
func FromPubKey(pubKey crypto.PubKey) (DID, error) {
	var code multicodec.Code

	switch pubKey.Type() {
	case pb.KeyType_Ed25519:
		code = multicodec.Ed25519Pub
	case pb.KeyType_RSA:
		code = RSA
	case pb.KeyType_Secp256k1:
		code = Secp256k1
	case pb.KeyType_ECDSA:
		var err error
		if code, err = codeForCurve(pubKey); err != nil {
			return Undef, err
		}
	default:
		return Undef, errors.New("unsupported key type")
	}

	if pubKey.Type() == pb.KeyType_ECDSA && code == Secp256k1 {
		var err error

		pubKey, err = coerceECDSAToSecp256k1(pubKey)
		if err != nil {
			return Undef, err
		}
	}

	var bytes []byte

	switch pubKey.Type() {
	case pb.KeyType_ECDSA:
		pkix, err := pubKey.Raw()
		if err != nil {
			return Undef, err
		}

		publicKey, err := x509.ParsePKIXPublicKey(pkix)
		if err != nil {
			return Undef, err
		}

		ecdsaPublicKey := publicKey.(*ecdsa.PublicKey)

		bytes = elliptic.MarshalCompressed(ecdsaPublicKey.Curve, ecdsaPublicKey.X, ecdsaPublicKey.Y)
	case pb.KeyType_Ed25519, pb.KeyType_Secp256k1:
		var err error

		if bytes, err = pubKey.Raw(); err != nil {
			return Undef, err
		}
	case pb.KeyType_RSA:
		var err error

		pkix, err := pubKey.Raw()
		if err != nil {
			return Undef, err
		}

		publicKey, err := x509.ParsePKIXPublicKey(pkix)
		if err != nil {
			return Undef, err
		}

		bytes = x509.MarshalPKCS1PublicKey(publicKey.(*rsa.PublicKey))
	}

	return DID{
		code:  code,
		bytes: string(append(varint.ToUvarint(uint64(code)), bytes...)),
	}, nil
}

// ToPubKey returns the crypto.PubKey encapsulated in the DID formed by
// parsing the provided string.
func ToPubKey(s string) (crypto.PubKey, error) {
	id, err := Parse(s)
	if err != nil {
		return nil, err
	}

	return id.PubKey()
}

func codeForCurve(pubKey crypto.PubKey) (multicodec.Code, error) {
	stdPub, err := crypto.PubKeyToStdKey(pubKey)
	if err != nil {
		return multicodec.Identity, err
	}

	ecdsaPub, ok := stdPub.(*ecdsa.PublicKey)
	if !ok {
		return multicodec.Identity, errors.New("failed to assert type for code to curve")
	}

	switch ecdsaPub.Curve {
	case elliptic.P256():
		return P256, nil
	case elliptic.P384():
		return P384, nil
	case elliptic.P521():
		return P521, nil
	case secp256k1.S256():
		return Secp256k1, nil
	default:
		return multicodec.Identity, fmt.Errorf("unsupported ECDSA curve: %s", ecdsaPub.Curve.Params().Name)
	}
}

// secp256k1.S256 is a valid ECDSA curve, but the go-libp2p/core/crypto
// package treats it as a different type and has a different format for
// the raw bytes of the public key.
//
// If a valid ECDSA public key was created using the secp256k1.S256 curve,
// this function will "convert" it from a crypto.ECDSAPubKey to a
// crypto.Secp256k1PublicKey.
func coerceECDSAToSecp256k1(pubKey crypto.PubKey) (crypto.PubKey, error) {
	stdPub, err := crypto.PubKeyToStdKey(pubKey)
	if err != nil {
		return nil, err
	}

	ecdsaPub, ok := stdPub.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to assert type for secp256k1 coersion")
	}

	ecdsaPubBytes := append([]byte{0x04}, append(ecdsaPub.X.Bytes(), ecdsaPub.Y.Bytes()...)...)

	secp256k1Pub, err := secp256k1.ParsePubKey(ecdsaPubBytes)
	if err != nil {
		return nil, err
	}

	cryptoPub := crypto.Secp256k1PublicKey(*secp256k1Pub)

	return &cryptoPub, nil
}
