//go:build wasm
// +build wasm

package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"syscall/js"

	"github.com/sonr-io/sonr/pkg/crypto/mpc"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
	txBetaAddress      = "v1-beta.sonr.ws:1317/cosmos/tx/v1beta/txs"
)

type motor struct {
	//	node *host.SonrHost
	wallet *mpc.Wallet
	// host   host.SonrHost
}

var instance *motor

// NewWallet creates a new mpc based wallet.
func NewWallet() error {
	if instance != nil {
		return errWalletExists
	}
	w, err := mpc.GenerateWallet()
	if err != nil {
		return err
	}
	instance = &motor{
		wallet: w,
	}
	return nil
}

// Address returns the address of the wallet.
func Address() (string, error) {
	if instance == nil {
		return "", nil
	}
	addr, err := instance.wallet.Address()
	if err != nil {
		return "", err
	}
	return addr, nil
}

// LoadWallet unmarshals the given JSON into the wallet.
func LoadWallet(buf []byte) error {
	if instance != nil {
		return errWalletExists
	}
	w := &mpc.MPCWallet{}
	err := w.Unmarshal(buf)
	if err != nil {
		return err
	}
	instance = &motor{
		wallet: w,
	}
	return nil
}

// MarshalWallet returns the JSON representation of the wallet.
func MarshalWallet() ([]byte, error) {
	if instance == nil {
		return nil, errWalletExists
	}
	buf, err := instance.wallet.Marshal()
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// Sign returns the signature of the given message.
func Sign(typeUrl string, msg []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	// tx, err := rt.UnmarshalTxMsg(typeUrl, msg)
	// if err != nil {
	// 	return nil, err
	// }

	sig, err := instance.wallet.Sign(msg)
	if err != nil {
		return nil, err
	}
	return mpc.SerializeSignature(sig)
}

// Verify returns true if the given signature is valid for the given message.
func Verify(msg []byte, sig []byte) bool {
	if instance == nil {
		return false
	}
	return instance.wallet.Verify(msg, sig)
}

// Broadcast broadcasts rawTx to the specified address
func Broadcast(addr string, tx []byte) error {
	res, err := http.Post(txBetaAddress, "application/json", bytes.NewBuffer(tx))
	if err != nil {
		return err
	}

	fmt.Print(res)
	return nil
}

/*
-------------------------------
		Wrappers
-------------------------------
*/
func NewWalletExport() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return NewWallet()
	})

	return js_func
}

func AddressExporter() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		addr, err := Address()

		if err != nil {
			return err
		}

		return addr
	})

	return js_func
}

func MarshalWalletWrapper() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		data, err := MarshalWallet()
		if err != nil {
			return err
		}

		return data
	})

	return js_func
}

func SignWrapper() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		url := args[0].String()
		msg := []byte(args[1].String())
		buf, err := Sign(url, msg)

		if err != nil {
			return err
		}

		return buf
	})

	return js_func
}

func VerifyWrapper() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		bytes := []byte(args[0].String())
		sig := []byte(args[1].String())
		res := Verify(bytes, sig)

		return res
	})

	return js_func
}

func BroadcastWrapper() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		addr := args[0].String()
		bytes := []byte(args[1].String())
		err := Broadcast(addr, bytes)

		return err
	})

	return js_func
}

func main() {
	fmt.Printf("Creating motor api")

	js.Global().Set("createWallet", NewWalletExport())
	js.Global().Set("getAddress", AddressExporter())
	js.Global().Set("marshalWallet", MarshalWalletWrapper())
	js.Global().Set("signTx", SignWrapper())
	js.Global().Set("verifyTx", VerifyWrapper())
	js.Global().Set("broadcastTx", BroadcastWrapper())

	fmt.Printf("Done creating motor api, methods available")

	// module cannot leave scope, keep the entry in scope
	<-make(chan (bool))
}
