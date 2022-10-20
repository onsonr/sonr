package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/third_party/types/common"
)

func (mtr *motorNodeImpl) triggerWalletEvent(event common.WalletEvent) error {
	if mtr.callback == nil {
		return fmt.Errorf("error callback is nil cannot trigger")
	}
	b, err := event.Marshal()

	if err != nil {
		return fmt.Errorf("Error while marshalling wallet event: \n%s", err.Error())
	}

	mtr.callback.OnWalletEvent(b)

	return nil
}
