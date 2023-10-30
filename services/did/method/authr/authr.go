package authr

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonr-io/kryptology/pkg/core/curves"
	"github.com/sonr-io/sonr/internal/crypto"
	"github.com/sonr-io/sonr/services/did/types"
)

const Method = types.DIDMethod("authr")

// Authenticator is a DID that can be used to authenticate a user
type Authenticator struct {
	Method types.DIDMethod
	ID     types.DIDIdentifier

	secret AuthSecretKey
}

// NewAuthenticator creates a new Authenticator DID
func NewAuthenticator(email string, cred *types.Credential, origin string) (*Authenticator, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := NewSecretKey()
	if err != nil {
		return nil, err
	}
	acc, err := key.AccumulatorKey()
	if err != nil {
		return nil, err
	}
	pub, err := acc.GetPublicKey(curve)
	if err != nil {
		return nil, err
	}
	pbz, err := pub.MarshalBinary()
	if err != nil {
		return nil, err
	}
	did := FormatIdentifier(pbz)
	m, id, err := types.ParseDID(did)
	if err != nil {
		return nil, err
	}
	credAuth, err := NewCredentialAuthData(cred, origin, key)
	if err != nil {
		return nil, err
	}
	ckey := credAuth.Key()
	credABz, err := cred.Serialize()
	if err != nil {
		return nil, err
	}
	id.AppendKeyList(ckey, crypto.Base64Encode(credABz))

	emailAuth, err := NewEmailAuthData(email, key)
	if err != nil {
		return nil, err
	}
	ekey := emailAuth.Key()
	emailABz, err := emailAuth.Marshal()
	if err != nil {
		return nil, err
	}
	id.AppendKeyList(ekey, crypto.Base64Encode(emailABz))
	m.SetKey(id.String(), "auth")
	return &Authenticator{
		Method: m,
		ID:     id,
	}, nil
}

// ResolveAuthenticator creates an Authenticator from a DID and an email.
func ResolveAuthenticator(didString string) (*Authenticator, error) {
	m, id, err := types.ParseDID(didString)
	if err != nil {
		return nil, err
	}
	return &Authenticator{
		Method: m,
		ID:     id,
	}, nil
}

// DID returns the DID string
func (a *Authenticator) DID() string {
	return a.DIDUrl().String()
}

// DIDSecretKey returns the DIDSecretKey for the authenticator
func (a *Authenticator) DIDSecretKey(email string) (types.DIDSecretKey, error) {
	vals := a.ID.GetKeyList("email")
	for _, val := range vals {
		bz, err := crypto.Base64Decode(val)
		if err != nil {
			return nil, err
		}
		auth := &AuthData{}
		err = auth.Unmarshal(bz)
		if err != nil {
			return nil, err
		}
		return auth.GetSecretKeyFromEmail(email)
	}
	return a.secret, nil
}

// DIDUrl returns the DIDUrl
func (a *Authenticator) DIDUrl() types.DIDUrl {
	return types.NewDIDUrl(a.Method, a.ID)
}

// ListCredentialDescriptors returns a list of credential descriptors for the authenticator
func (a *Authenticator) ListCredentialDescriptors(origin string) ([]protocol.CredentialDescriptor, error) {
	creds := make([]protocol.CredentialDescriptor, 0)
	for _, credStr := range a.ID.GetKeyList(fmt.Sprintf("credentials/%s", origin)) {
		credBz, err := crypto.Base64Decode(credStr)
		if err != nil {
			return nil, err
		}
		var cred types.Credential
		err = cred.Deserialize([]byte(credBz))
		if err != nil {
			return nil, err
		}
		creds = append(creds, cred.GetDescriptor())
	}
	return creds, nil
}

// LockResource locks a resource to the authenticator
func (a *Authenticator) LockResource(resource types.DIDResource) ([]byte, error) {
	bz, err := resource.Data()
	if err != nil {
		return nil, err
	}
	encBz, err := a.secret.Encrypt(bz)
	if err != nil {
		return nil, err
	}
	err = resource.Update(encBz)
	if err != nil {
		return nil, err
	}
	return encBz, nil
}

// UnlockResource unlocks a resource from the authenticator
func (a *Authenticator) UnlockResource(resource types.DIDResource) ([]byte, error) {
	bz, err := resource.Data()
	if err != nil {
		return nil, err
	}
	decBz, err := a.secret.Decrypt(bz)
	if err != nil {
		return nil, err
	}
	err = resource.Update(decBz)
	if err != nil {
		return nil, err
	}
	return decBz, nil
}
