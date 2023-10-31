package didcontroller

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonr-io/kryptology/pkg/core/curves"

	"github.com/sonrhq/sonr/internal/crypto"
	"github.com/sonrhq/sonr/pkg/didauthz"
	"github.com/sonrhq/sonr/pkg/didcommon"
)

const Method = didcommon.Method("authr")

// Authenticator is a DID that can be used to authenticate a user
type Authenticator struct {
	Method didcommon.Method
	ID     didcommon.Identifier

	secret didauthz.AuthSecretKey
}

// NewAuthenticator creates a new Authenticator DID
func NewAuthenticator(email string, cred *didcommon.Credential, origin string) (*Authenticator, error) {
	curve := curves.BLS12381(&curves.PointBls12381G1{})
	key, err := didauthz.NewSecretKey()
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
	did := didauthz.FormatIdentifier(pbz)
	m, id, err := didcommon.Parse(did)
	if err != nil {
		return nil, err
	}
	credAuth, err := didauthz.NewCredentialAuthData(cred, origin, key)
	if err != nil {
		return nil, err
	}
	ckey := credAuth.Key()
	credABz, err := cred.Serialize()
	if err != nil {
		return nil, err
	}
	id.AppendKeyList(ckey, crypto.Base64Encode(credABz))

	emailAuth, err := didauthz.NewEmailAuthData(email, key)
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
	m, id, err := didcommon.Parse(didString)
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
func (a *Authenticator) DIDSecretKey(email string) (didcommon.SecretKey, error) {
	vals := a.ID.GetKeyList("email")
	for _, val := range vals {
		bz, err := crypto.Base64Decode(val)
		if err != nil {
			return nil, err
		}
		auth := &didauthz.AuthData{}
		err = auth.Unmarshal(bz)
		if err != nil {
			return nil, err
		}
		return auth.GetSecretKeyFromEmail(email)
	}
	return a.secret, nil
}

// DIDUrl returns the DIDUrl
func (a *Authenticator) DIDUrl() didcommon.Url {
	return didcommon.NewUrl(a.Method, a.ID)
}

// ListCredentialDescriptors returns a list of credential descriptors for the authenticator
func (a *Authenticator) ListCredentialDescriptors(origin string) ([]protocol.CredentialDescriptor, error) {
	creds := make([]protocol.CredentialDescriptor, 0)
	for _, credStr := range a.ID.GetKeyList(fmt.Sprintf("credentials/%s", origin)) {
		credBz, err := crypto.Base64Decode(credStr)
		if err != nil {
			return nil, err
		}
		var cred didcommon.Credential
		err = cred.Deserialize([]byte(credBz))
		if err != nil {
			return nil, err
		}
		creds = append(creds, cred.GetDescriptor())
	}
	return creds, nil
}

// LockResource locks a resource to the authenticator
func (a *Authenticator) LockResource(resource didcommon.DIDResource) ([]byte, error) {
	bz, err := a.ID.FetchResource(resource.Key())
	if err != nil {
		return nil, err
	}
	encBz, err := a.secret.Encrypt(bz)
	if err != nil {
		return nil, err
	}
	_, err = a.ID.AddResource(resource.Key(), encBz)
	if err != nil {
		return nil, err
	}
	return encBz, nil
}

// UnlockResource unlocks a resource from the authenticator
func (a *Authenticator) UnlockResource(resource didcommon.DIDResource) ([]byte, error) {
	bz, err := a.ID.FetchResource(resource.Key())
	if err != nil {
		return nil, err
	}
	decBz, err := a.secret.Decrypt(bz)
	if err != nil {
		return nil, err
	}
	_, err = a.ID.AddResource(resource.Key(), decBz)
	if err != nil {
		return nil, err
	}
	return decBz, nil
}
