package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/jwt"
	"github.com/spf13/viper"

	"github.com/sonr-io/sonr/internal/crypto"
	identitytypes "github.com/sonr-io/sonr/x/identity/types"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 Credential JWT                                 ||
// ! ||--------------------------------------------------------------------------------||

// CredentialJWTClaims struct is defining a data structure that represents the claims (or payload) of a JSON Web Token (JWT). It contains fields such as `Did`, `ExpiresAt`, and `UUID`, which hold information related to the JWT. These fields can be used to store additional data that
// needs to be included in the JWT for authentication or authorization purposes.
type CredentialJWTClaims struct {
	AccountDid string `json:"did"`
	ExpiresAt  int64  `json:"expires_at"`
	UUID       string `json:"uuid"`
}

// NewCredentialJWTClaims takes a did string and returns a JWTClaims struct.
func NewCredentialJWTClaims(did string) (CredentialJWTClaims, string, error) {
	claims := CredentialJWTClaims{
		AccountDid: did,
		ExpiresAt:  time.Now().Add(time.Hour).Unix(),
		UUID:       uuid.NewString(),
	}
	token, err := jwt.Sign(jwt.HS256, envJWTSigningKey(), claims)
	if err != nil {
		return CredentialJWTClaims{}, "", err
	}
	return claims, crypto.Base58Encode(token), nil
}

// VerifyCredentialJWTClaims takes a token string as input and returns the JWTClaims and an error.
func VerifyCredentialJWTClaims(token string) (CredentialJWTClaims, error) {
	raw, err := crypto.Base58Decode(token)
	if err != nil {
		return CredentialJWTClaims{}, err
	}
	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, envJWTSigningKey(), raw)
	if err != nil {
		return CredentialJWTClaims{}, err
	}
	// Extract the claims
	var claims CredentialJWTClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return CredentialJWTClaims{}, err
	}
	return claims, nil
}

// IsValid returns true if the credential is valid
func (c CredentialJWTClaims) IsValid() error {
	if time.Now().Unix() > c.ExpiresAt {
		return ErrJWTExpired
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                    Email JWT                                   ||
// ! ||--------------------------------------------------------------------------------||

// EmailJWTClaims struct is defining a data structure that represents the claims (or payload) of a JSON Web Token (JWT). It contains fields such as `Did`, `ExpiresAt`, and `UUID`, which hold information related to the JWT. These fields can be used to store additional data that
// needs to be included in the JWT for authentication or authorization purposes.
type EmailJWTClaims struct {
	AccountDid    string `json:"did"`
	ControllerDid string `json:"controller_did"`
	ExpiresAt     int64  `json:"expires_at"`
	Email         string `json:"email"`
}

// NewEmailJWTClaims takes a did string and returns a JWTClaims struct.
func NewEmailJWTClaims(did string, email string) (EmailJWTClaims, string, error) {
	claims := EmailJWTClaims{
		AccountDid: did,
		ExpiresAt:  time.Now().Add(time.Minute * 5).Unix(),
		Email:      email,
	}
	token, err := jwt.Sign(jwt.HS256, envJWTSigningKey(), claims)
	if err != nil {
		return EmailJWTClaims{}, "", err
	}
	return claims, crypto.Base58Encode(token), nil
}

// VerifyEmailJWTClaims takes a token string as input and returns the JWTClaims and an error.
func VerifyEmailJWTClaims(token string) (EmailJWTClaims, error) {
	raw, err := crypto.Base58Decode(token)
	if err != nil {
		return EmailJWTClaims{}, err
	}

	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, envJWTSigningKey(), raw)
	if err != nil {
		return EmailJWTClaims{}, err
	}
	// Extract the claims
	var claims EmailJWTClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return EmailJWTClaims{}, err
	}
	if time.Now().Unix() > claims.ExpiresAt {
		return EmailJWTClaims{}, ErrJWTExpired
	}
	return claims, nil
}

// IsValid returns true if the credential is valid
func (c EmailJWTClaims) IsValid() error {
	if time.Now().Unix() > c.ExpiresAt {
		return ErrJWTExpired
	}
	return nil
}

// ControllerIdentifier returns the last part of the DID used to anonymize email addresses
func (c EmailJWTClaims) ControllerIdentifier() string {
	ptrs := strings.Split(c.ControllerDid, ":")
	return ptrs[len(ptrs)-1]
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                             Email-based Encryption                             ||
// ! ||--------------------------------------------------------------------------------||

// EmailEncryptedData struct is defining a data structure that represents the encrypted data to be stored in the JWT. It contains fields such as `Did` and `Base64`, which hold information related to the encrypted data. The `Did` field represents the decentralized identifier
// associated with the data, and the `Base64` field stores the encrypted data in base64 format. This struct is used in the `Encrypt` and `Decrypt` methods of the `EmailJWTClaims` struct to encrypt and decrypt data using the JWT signing key.
type EmailEncryptedData struct {
	Did    string `json:"did"`
	Base64 string `json:"base64"`
}

// Encrypt method of the `EmailJWTClaims` struct is used to encrypt data using the JWT signing key. It takes a byte slice `data` as input and returns a byte slice representing the encrypted data and an error.
func (c EmailJWTClaims) Encrypt(data []byte) ([]byte, error) {
	enc := EmailEncryptedData{
		Did:    c.AccountDid,
		Base64: crypto.Base64Encode(data),
	}
	token, err := jwt.Sign(jwt.HS256, c.SigningKey(), enc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Decrypt method of the `EmailJWTClaims` struct is used to decrypt data that has been encrypted using the JWT signing key. It takes a byte slice `data` as input, which represents the encrypted data, and returns a byte slice representing the decrypted data and an error.
func (c EmailJWTClaims) Decrypt(data []byte) ([]byte, error) {
	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, c.SigningKey(), data)
	if err != nil {
		return nil, err
	}
	// Extract the claims
	var enc EmailEncryptedData
	err = verifiedToken.Claims(&enc)
	if err != nil {
		return nil, err
	}
	return crypto.Base64Decode(enc.Base64)
}

// SigningKey returns the JWT signing key
func (c EmailJWTClaims) SigningKey() []byte {
	ptrs := strings.Split(c.AccountDid, ":")
	id := ptrs[len(ptrs)-1]
	str := fmt.Sprintf("%s+%s", c.Email, id)
	return []byte(str)
}

// ! ||-----------------------------------------------------------------------------||
// ! ||                                 Session JWT                                 ||
// ! ||-----------------------------------------------------------------------------||

// SessionJWTClaims struct is defining a data structure that represents the claims (or payload) of a JSON Web Token (JWT). It contains fields such as `Did`, `ExpiresAt`, and `UUID`, which hold information related to the JWT. These fields can be used to store additional data that
// needs to be included in the JWT for authentication or authorization purposes.
type SessionJWTClaims struct {
	Address       string `json:"address"`
	Authenticator string `json:"authenticator"`
	Email         string `json:"email"`
	ExpiresAt     int64  `json:"expires_at"`
}

// NewSessionJWTClaims takes a did string and returns a JWTClaims struct.
func NewSessionJWTClaims(email string, acc *identitytypes.ControllerAccount) (string, error) {
	claims := SessionJWTClaims{
		Address:       acc.Address,
		Authenticator: acc.Authenticators[0],
		Email:         email,
		ExpiresAt:     time.Now().Add(time.Minute * 30).Unix(),
	}
	token, err := jwt.Sign(jwt.HS256, envJWTSigningKey(), claims)
	if err != nil {
		return "", err
	}
	return crypto.Base58Encode(token), nil
}

// VerifySessionJWTClaims takes a token string as input and returns the JWTClaims and an error.
func VerifySessionJWTClaims(token string) (SessionJWTClaims, error) {
	raw, err := crypto.Base58Decode(token)
	if err != nil {
		return SessionJWTClaims{}, err
	}
	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, envJWTSigningKey(), raw)
	if err != nil {
		return SessionJWTClaims{}, err
	}
	// Extract the claims
	var claims SessionJWTClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return SessionJWTClaims{}, err
	}
	return claims, nil
}

// IsValid returns true if the credential is valid
func (c SessionJWTClaims) IsValid() error {
	if time.Now().Unix() > c.ExpiresAt {
		return ErrJWTExpired
	}
	return nil
}

// envJWTSigningKey returns the JWT signing key
func envJWTSigningKey() []byte {
	return []byte(viper.GetString("highway.jwt.key"))
}
