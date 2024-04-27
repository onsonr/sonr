package types

import (
	"regexp"

	"github.com/cosmos/btcutil/bech32"
	"github.com/ethereum/go-ethereum/crypto"
)

// BTCAddress is a type for the BTC address
type BTCAddress string

// Bytes returns the bytes representation of the BTC address
func (a BTCAddress) Bytes() []byte {
	return []byte(a)
}

// DID returns the DID representation of the BTC address given a method
func (a BTCAddress) DID(method string) string {
	return "did:" + method + ":" + a.String()
}

// String returns the string representation of the BTC address
func (a BTCAddress) String() string {
	return string(a)
}

// Validate returns an error if the BTC address is invalid
func (a BTCAddress) Validate() error {
	re := regexp.MustCompile(`\b(bc(0([ac-hj-np-z02-9]{39}|[ac-hj-np-z02-9]{59})|1[ac-hj-np-z02-9]{8,87})|[13][a-km-zA-HJ-NP-Z1-9]{25,35})\b`)
	if !re.MatchString(string(a)) {
		return ErrInvalidBTCAddressFormat
	}
	return nil
}

// ETHAddress is a type for the ETH address
type ETHAddress string

// Bytes returns the bytes representation of the ETH address
func (a ETHAddress) Bytes() []byte {
	return []byte(a)
}

// DID returns the DID representation of the ETH address given a method
func (a ETHAddress) DID(method string) string {
	return "did:" + method + ":" + a.String()
}

// String returns the string representation of the ETH address
func (a ETHAddress) String() string {
	return string(a)
}

// Validate returns an error if the ETH address is invalid
func (a ETHAddress) Validate() error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(string(a)) {
		return ErrInvalidETHAddressFormat
	}
	return nil
}

// IDXAddress is a type for the IDX address
type IDXAddress string

// Bytes returns the bytes representation of the IDX address
func (a IDXAddress) Bytes() []byte {
	return []byte(a)
}

// DID returns the DID representation of the IDX address given a method
func (a IDXAddress) DID(method string) string {
	return "did:" + method + ":" + a.String()
}

// String returns the string representation of the IDX address
func (a IDXAddress) String() string {
	return string(a)
}

// GetBTCAddress returns the BTC address from the public key using bech32 encoding
func GetBTCAddress(publicKey *PublicKey) (BTCAddress, error) {
	addr, err := bech32.Encode("bc", publicKey.Bytes())
	if err != nil {
		return "", err
	}
	return BTCAddress(addr), nil
}

// GetETHAddress returns the ETH address from the public key using keccak256
func GetETHAddress(publicKey *PublicKey) (ETHAddress, error) {
	ecdsaPub, err := crypto.DecompressPubkey(publicKey.Bytes())
	if err != nil {
		return "", err
	}
	addr := crypto.PubkeyToAddress(*ecdsaPub).Hex()
	return ETHAddress(addr), nil
}

// GetIDXAddress returns the IDX address from the public key
func GetIDXAddress(publicKey *PublicKey) (IDXAddress, error) {
	addr, err := bech32.Encode("idx", publicKey.Bytes())
	if err != nil {
		return "", err
	}
	return IDXAddress(addr), nil
}
