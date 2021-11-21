package exchange

import (
	"github.com/hyperledger/aries-framework-go/pkg/client/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	"github.com/sonr-io/core/wallet"
)

func NewDIDSwap() error {
	bob, err := didexchange.New(wallet.Provider)
	if err != nil {
		return err
	}

	bobActions := make(chan service.DIDCommAction, 1)
	err = bob.RegisterActionEvent(bobActions)
	if err != nil {
		return err
	}

	go func() {
		service.AutoExecuteActionEvent(bobActions)
	}()
	return nil
}
