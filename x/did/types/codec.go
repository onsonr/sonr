package types

import (
	"crypto/ecdsa"
	"strings"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

var (
	amino    = codec.NewLegacyAmino()
	AminoCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateParams{}, ModuleName+"/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*cryptotypes.PubKey)(nil),
		//		&PubKey{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgRegisterController{},
		&MsgRegisterService{},
		&MsgAllocateVault{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func ComputeEthAddress(pk ecdsa.PublicKey) string {
	// Generate Ethereum address
	address := ethcrypto.PubkeyToAddress(pk)

	// Apply ERC-55 checksum encoding
	addr := address.Hex()
	addr = strings.ToLower(addr)
	addr = strings.TrimPrefix(addr, "0x")
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(addr))
	hashBytes := hash.Sum(nil)

	result := "0x"
	for i, c := range addr {
		if c >= '0' && c <= '9' {
			result += string(c)
		} else {
			if hashBytes[i/2]>>(4-i%2*4)&0xf >= 8 {
				result += strings.ToUpper(string(c))
			} else {
				result += string(c)
			}
		}
	}
	return result
}
