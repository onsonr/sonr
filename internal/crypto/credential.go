package crypto

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/libp2p/go-libp2p/core/peer"
	types "github.com/sonrhq/core/types/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// PartyID is a type alias for party.ID in pkg/party.
type PartyID = party.ID

// PeerID is a type alias for peer.ID in pkg/peer.
type PeerID = peer.ID

// MPCPool is a type alias for pool.Pool in pkg/pool.
type MPCPool = pool.Pool

// MPCCmpConfig is a type alias for pool.CmpConfig in pkg/pool.
type MPCCmpConfig = cmp.Config

// MPCECDSASignature is a type alias for ecdsa.Signature in pkg/ecdsa.
type MPCECDSASignature = ecdsa.Signature

// MPCSecp256k1Curve is a type alias for curve.Secp256k1Point in pkg/math/curve.
type MPCSecp256k1Curve = curve.Secp256k1

// MPCSecp256k1Point is a type alias for curve.Secp256k1Point in pkg/math/curve.
type MPCSecp256k1Point = curve.Secp256k1Point

// NewMPCPool creates a new MPCPool with the given size.
func NewMPCPool(size int) *MPCPool {
	return pool.NewPool(0)
}

// NewEmptyECDSASecp256k1Signature creates a new empty MPCECDSASignature.
func NewEmptyECDSASecp256k1Signature() MPCECDSASignature {
	return ecdsa.EmptySignature(MPCSecp256k1Curve{})
}

// NewWebAuthnCredential creates a new WebauthnCredential from a ParsedCredentialCreationData and contains all needed information about a WebAuthn credential for storage.
// This is then used to create a VerificationMethod for the DID Document.
func NewWebAuthnCredential(c *protocol.ParsedCredentialCreationData) *types.WebauthnCredential {
	transportsStr := []string{}
	return &types.WebauthnCredential{
		Id:              c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		AttestationType: c.Response.AttestationObject.Format,
		Transport:       transportsStr,
		Authenticator: &types.WebauthnAuthenticator{
			Aaguid:    c.Response.AttestationObject.AuthData.AttData.AAGUID,
			SignCount: c.Response.AttestationObject.AuthData.Counter,
		},
	}
}
