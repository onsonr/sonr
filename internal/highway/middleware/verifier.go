package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"

	"github.com/sonrhq/core/internal/highway/types"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/did/controller"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

// The function GetAuthCookies takes a gin.Context as input and returns three strings and an error.
func fetchAuthCookies(c *gin.Context) (string, string, string, error) {
	jwtToken, err := c.Cookie("sonr-jwt")
	if err != nil {
		return "", "", "", fmt.Errorf("no jwt cookie found")
	}
	did, err := c.Cookie("sonr-did")
	if err != nil {
		return "", "", "", fmt.Errorf("no did cookie found")
	}
	alias, err := c.Cookie("sonr-alias")
	if err != nil {
		return "", "", "", fmt.Errorf("no alias cookie found")
	}
	return jwtToken, did, alias, nil
}

// StoreAuthCookies function stores authentication cookies in the context.
func StoreAuthCookies(c *gin.Context, res *types.AuthenticationResult, origin string) gin.H {
	c.SetCookie("sonr-jwt", res.JWT, 1800, "/", origin, true, true)
	c.SetCookie("sonr-did", res.DID, 1800, "/", origin, true, false)
	c.SetCookie("sonr-alias", res.Alias, 1800, "/", origin, true, false)
	return gin.H{
		"success":      true,
		"account":      res.Account,
		"did_document": res.DIDDocument,
		"token":        res.JWT,
	}
}

// IssueCredentialAttestationOptions takes a ucwId alias, and random unclaimed address and returns a token with the credential options.
func IssueCredentialAttestationOptions(alias string, record *servicetypes.ServiceRecord) (string, protocol.URLEncodedBase64, error) {
	chal, err := crypto.GenerateChallenge()
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate challenge: %w", err)
	}
	attestionOpts, err := record.GetCredentialCreationOptions(alias, chal)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get credential creation options: %w", err)
	}
	return attestionOpts, chal, nil
}

// IssueCredentialAssertionOptions takes a didDocument and serviceRecord in order to create a credential options.
func IssueCredentialAssertionOptions(email string, record *servicetypes.ServiceRecord) (string, protocol.URLEncodedBase64, string, error) {
	addr, err := GetEmailRecordCreator(email)
	if err != nil {
		return "", nil, "", fmt.Errorf("failed to get email record creator: %w", err)
	}
	controllerAcc, err := GetControllerAccount(addr)
	if err != nil {
		return "", nil, "", fmt.Errorf("failed to get controller account: %w", err)
	}
	cont, err := controller.AuthorizeIdentity(email, controllerAcc)
	if err != nil {
		return "", nil, "", fmt.Errorf("failed to authorize identity: %w", err)
	}
	vms, err := cont.Authenticator.ListCredentialDescriptors(record.GetBaseOrigin())
	chal, err := crypto.GenerateChallenge()
	if err != nil {
		return "", nil, "", fmt.Errorf("failed to generate challenge: %w", err)
	}
	assertionOpts, err := record.GetCredentialAssertionOptions(vms, chal)
	if err != nil {
		return "", nil, "", fmt.Errorf("failed to get credential assertion options: %w", err)
	}
	return assertionOpts, chal, cont.Account().Address, nil
}

// IssueEmailAssertionOptions takes a didDocument and serviceRecord in order to create a credential options.
func IssueEmailAssertionOptions(email string, ucwDid string) (string, error) {
	_, tkn, err := types.NewEmailJWTClaims(ucwDid, email)
	if err != nil {
		return "", fmt.Errorf("failed to create email claims: %w", err)
	}
	return tkn, nil
}

// UseControllerAccount takes a jwt token and returns a controller account.
func UseControllerAccount(token string) (*controller.SonrController, error) {
	claims, err := types.VerifySessionJWTClaims(token)
	if err != nil {
		return nil, fmt.Errorf("failed to verify claims: %w", err)
	}
	acc, err := GetControllerAccount(claims.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get controller account: %w", err)
	}
	controller, err := controller.AuthorizeIdentity(claims.Email, acc)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize identity: %w", err)
	}
	return controller, nil
}
