package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/x/registry/types"
)

func CreateAccount() error {
	c := client.NewClient(true)
	w, err := crypto.Generate()
	if err != nil {
		return err
	}
	bechAddr, err := w.Bech32Address()
	if err != nil {
		return err
	}
	err = c.RequestFaucet(bechAddr)
	if err != nil {
		return err
	}

	acc, err := c.QueryAccount(bechAddr)
	if err != nil {
		return err
	}

	doc, err := w.DIDDocument()
	if err != nil {
		return err
	}

	docBz, err := doc.MarshalJSON()
	if err != nil {
		return err
	}
	msg1 := types.NewMsgCreateWhoIs(bechAddr, docBz, types.WhoIsType_USER)
	ai, txb, err := crypto.BuildTx(w, msg1)
	if err != nil {
		return err
	}

	sig, err := w.SignTx(acc, ai, txb)
	if err != nil {
		return err
	}
	resp, err := c.BroadcastTx(txb, sig, ai)
	if err != nil {
		return err
	}
	fmt.Println(resp.String())
	return nil
}
