//go:build wasm
// +build wasm

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"syscall/js"

	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

type motor struct {
	//	node *host.SonrHost
	wallet *crypto.MPCWallet
	// host   host.SonrHost
}

var instance *motor

// NewWallet creates a new mpc based wallet.
func NewWallet() error {
	if instance != nil {
		return errWalletExists
	}
	w, err := crypto.GenerateWallet()
	if err != nil {
		return err
	}
	instance = &motor{
		wallet: w,
	}
	return nil
}

// Address returns the address of the wallet.
func Address() string {
	if instance == nil {
		return ""
	}
	addr, err := instance.wallet.Address()
	if err != nil {
		return ""
	}
	return addr
}

// LoadWallet unmarshals the given JSON into the wallet.
func LoadWallet(buf []byte) error {
	if instance != nil {
		return errWalletExists
	}
	w := &crypto.MPCWallet{}
	err := w.Unmarshal(buf)
	if err != nil {
		return err
	}
	instance = &motor{
		wallet: w,
	}
	return nil
}

// DidDoc returns the DID document as JSON
func DidDoc() string {
	if instance == nil {
		return ""
	}
	buf, err := instance.wallet.DIDDocument.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

// ImportCredentials imports the given credentials into the wallet.
func ImportCredential(buf []byte) error {
	if instance == nil {
		return errWalletNotExists
	}
	var cred did.Credential
	err := json.Unmarshal(buf, &cred)
	if err != nil {
		return err
	}
	vmdid, err := did.ParseDID(fmt.Sprintf("%s#%s", instance.wallet.DID, cred.ID))
	if err != nil {
		return err
	}
	vm := &did.VerificationMethod{
		ID:         *vmdid,
		Type:       ssi.ECDSASECP256K1VerificationKey2019,
		Controller: instance.wallet.DID,
		Credential: &cred,
	}
	instance.wallet.DIDDocument.AddAssertionMethod(vm)
	return nil
}

// MarshalWallet returns the JSON representation of the wallet.
func MarshalWallet() []byte {
	if instance == nil {
		return nil
	}
	buf, err := instance.wallet.Marshal()
	if err != nil {
		return nil
	}
	return buf
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
	return crypto.SerializeSignature(sig)
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
	apiEndpoint := "v1-beta.sonr.ws:1317/cosmos/tx/v1beta/txs"
	res, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(tx))
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
		return Address()
	})

	return js_func
}

func DidDocExporter() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return DidDoc()
	})

	return js_func
}

func ImportCredentialWrapper() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		buf := []byte(args[0].String())
		return ImportCredential(buf)
	})

	return js_func
}

func MarshalWalletWrapper() js.Func {
	js_func := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return MarshalWallet()
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
	js.Global().Set("getDidDoc", DidDocExporter())
	js.Global().Set("importCredential", ImportCredentialWrapper())
	js.Global().Set("marshalWallet", MarshalWalletWrapper())
	js.Global().Set("signTx", SignWrapper())
	js.Global().Set("verifyTx", VerifyWrapper())
	js.Global().Set("broadcastTx", BroadcastWrapper())

	fmt.Printf("Done creating motor api, methods available")

	// module cannot leave scope, keep the entry in scope
	<-make(chan (bool))
}
