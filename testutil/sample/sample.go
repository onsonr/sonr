package sample

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// "Read the file named `name` and return the contents as a byte slice. If there's an error, panic."
//
// The `ioutil.ReadFile` function returns two values: the data and an error. If there's an error, we
// panic. If there's no error, we return the data
func ReadTestFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return data
}
