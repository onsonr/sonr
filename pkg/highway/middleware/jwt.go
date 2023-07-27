package middleware

import (
	"context"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/highlight/highlight/sdk/highlight-go"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/pkg/did/controller"
	types "github.com/sonrhq/core/pkg/highway/types"

	// "github.com/sonrhq/core/pkg/sfs/store"
	domaintypes "github.com/sonrhq/core/x/domain/types"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

type ClaimsAPI struct{}

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
func IssueCredentialAssertionOptions(email string, record *servicetypes.ServiceRecord) (string, protocol.URLEncodedBase64, error) {
	addr, err := GetEmailRecordCreator(email)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get email record creator: %w", err)
	}
	controllerAcc, err := GetControllerAccount(addr)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get controller account: %w", err)
	}
	cont, err := controller.AuthorizeIdentity(email, controllerAcc)
	if err != nil {
		return "", nil, fmt.Errorf("failed to authorize identity: %w", err)
	}
	vms, err := cont.Authenticator.ListCredentialDescriptors(record.GetBaseOrigin())
	chal, err := crypto.GenerateChallenge()
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate challenge: %w", err)
	}
	assertionOpts, err := record.GetCredentialAssertionOptions(vms, chal)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get credential assertion options: %w", err)
	}
	return assertionOpts, chal, nil
}

// IssueEmailAssertionOptions takes a didDocument and serviceRecord in order to create a credential options.
func IssueEmailAssertionOptions(email string, ucwDid string) (string, error) {
	_, tkn, err := types.NewEmailJWTClaims(ucwDid, email)
	if err != nil {
		return "", fmt.Errorf("failed to create email claims: %w", err)
	}
	return tkn, nil
}

func PublishControllerAccount(alias string, cred *servicetypes.WebauthnCredential, origin string) (*controller.SonrController, *types.TxResponse, error) {
	ctx := context.Background()
	controller, err := controller.New(alias, cred, origin)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create controller: %w", err)

	}
	acc := controller.Account()
	accMsg := identitytypes.NewMsgCreateControllerAccount(acc.Address, acc.PublicKey, acc.Authenticators...)
	usrMsg := domaintypes.NewMsgCreateEmailUsernameRecord(acc.Address, alias)
	resp, err := controller.GetPrimaryWallet().SendTx(accMsg, usrMsg)
	if err != nil {
		highlight.RecordError(ctx, err)
		return nil, nil, fmt.Errorf("failed to send tx: %w", err)
	}
	fmt.Println(resp)
	return controller, resp, nil
}

func UseControllerAccount(token string) (*controller.SonrController, error) {
	ctx := context.Background()
	claims, err := types.VerifySessionJWTClaims(token)
	if err != nil {
		highlight.RecordError(ctx, err)
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
