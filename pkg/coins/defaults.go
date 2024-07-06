package coins

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

type coin struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Hrp    string `json:"hrp"`
	Method string `json:"method"`
	Index  int64  `json:"index"`
	Path   uint32 `json:"path"`
}

// FormatAddress formats the address based on the coin
func (c *coin) FormatAddress(pubKey []byte) (string, error) {
	if c.Hrp != "" {
		return bech32.ConvertAndEncode(c.Hrp, pubKey)
	}
	return "", fmt.Errorf("unsupported coin")
}

// GetIndex returns the coin index
func (c *coin) GetIndex() int64 {
	return c.Index
}

// GetName returns the coin name
func (c *coin) GetName() string {
	return c.Name
}

// GetSymbol returns the coin symbol
func (c *coin) GetSymbol() string {
	return c.Symbol
}

// GetHrp returns the coin hrp
func (c *coin) GetHrp() string {
	return c.Hrp
}

// GetPath returns the coin path
func (c *coin) GetPath() uint32 {
	return c.Path
}

// GetMethod returns the DID method for the coin
func (c *coin) GetMethod() string {
	return c.Method
}
