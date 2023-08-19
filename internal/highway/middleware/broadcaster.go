package middleware

import (
	"fmt"

	types "github.com/sonrhq/core/internal/highway/types"
	"github.com/sonrhq/core/pkg/did/controller"

	// "github.com/sonrhq/core/pkg/sfs/store"
	domaintypes "github.com/sonrhq/core/x/domain/types"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

// PublishControllerAccount creates a new controller account and publishes it to the blockchain
func PublishControllerAccount(alias string, cred *servicetypes.WebauthnCredential, origin string) (*controller.SonrController, *types.TxResponse, error) {
	controller, err := controller.New(alias, cred, origin)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create controller: %w", err)

	}
	acc := controller.Account()
	accMsg := identitytypes.NewMsgCreateControllerAccount(acc.Address, acc.PublicKey, acc.Authenticators...)
	usrMsg := domaintypes.NewMsgCreateEmailUsernameRecord(acc.Address, alias)
	resp, err := controller.GetPrimaryWallet().SendTx(accMsg, usrMsg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send tx: %w", err)
	}
	fmt.Println(resp)
	return controller, resp, nil
}

// CreateOrganizationRecord creates a new organization record and publishes it to the blockchain
func CreateOrganizationRecord(name string, origin string, admin string, controller *controller.SonrController) error {
	return nil
}
