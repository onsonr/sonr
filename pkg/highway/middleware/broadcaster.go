package middleware

import (
	"context"
	"fmt"

	"github.com/highlight/highlight/sdk/highlight-go"
	"github.com/sonrhq/core/pkg/did/controller"
	types "github.com/sonrhq/core/pkg/highway/types"

	// "github.com/sonrhq/core/pkg/sfs/store"
	domaintypes "github.com/sonrhq/core/x/domain/types"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

type ClaimsAPI struct{}


func PublishControllerAccount(alias string, cred *servicetypes.WebauthnCredential, origin string) (*controller.SonrController, *types.TxResponse, error) {
	ctx := context.Background()
	controller, err := controller.New(alias, cred, origin)
	if err != nil {
		highlight.RecordError(ctx, err)
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

