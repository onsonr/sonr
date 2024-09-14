package types

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	fmt "fmt"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/mr-tron/base58/base58"
	"github.com/onsonr/crypto"
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
		&PubKey{},
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

type ChainCode uint32

const (
	ChainCodeBTC ChainCode = 0
	ChainCodeETH ChainCode = 60
	ChainCodeIBC ChainCode = 118
	ChainCodeSNR ChainCode = 703
)

var InitialChainCodes = map[DIDNamespace]ChainCode{
	DIDNamespace_DID_NAMESPACE_BITCOIN:  ChainCodeBTC,
	DIDNamespace_DID_NAMESPACE_IBC:      ChainCodeIBC,
	DIDNamespace_DID_NAMESPACE_ETHEREUM: ChainCodeETH,
	DIDNamespace_DID_NAMESPACE_SONR:     ChainCodeSNR,
}

func (c ChainCode) FormatAddress(pubKey *PubKey) (string, error) {
	switch c {
	case ChainCodeBTC:
		return bech32.Encode("bc", pubKey.Bytes())

	case ChainCodeETH:
		epk, err := pubKey.ECDSA()
		if err != nil {
			return "", err
		}
		return ComputeEthAddress(*epk), nil

	case ChainCodeSNR:
		return bech32.Encode("idx", pubKey.Bytes())

	case ChainCodeIBC:
		return bech32.Encode("cosmos", pubKey.Bytes())

	}
	return "", ErrUnsopportedChainCode
}

func (n DIDNamespace) ChainCode() (uint32, error) {
	switch n {
	case DIDNamespace_DID_NAMESPACE_BITCOIN:
		return 0, nil
	case DIDNamespace_DID_NAMESPACE_ETHEREUM:
		return 64, nil
	case DIDNamespace_DID_NAMESPACE_IBC:
		return 118, nil
	case DIDNamespace_DID_NAMESPACE_SONR:
		return 703, nil
	default:
		return 0, fmt.Errorf("unsupported chain")
	}
}

func (n DIDNamespace) DIDMethod() string {
	switch n {
	case DIDNamespace_DID_NAMESPACE_IPFS:
		return "ipfs"
	case DIDNamespace_DID_NAMESPACE_SONR:
		return "sonr"
	case DIDNamespace_DID_NAMESPACE_BITCOIN:
		return "btcr"
	case DIDNamespace_DID_NAMESPACE_ETHEREUM:
		return "ethr"
	case DIDNamespace_DID_NAMESPACE_IBC:
		return "ibcr"
	case DIDNamespace_DID_NAMESPACE_WEBAUTHN:
		return "webauthn"
	case DIDNamespace_DID_NAMESPACE_DWN:
		return "motr"
	case DIDNamespace_DID_NAMESPACE_SERVICE:
		return "web"
	default:
		return "n/a"
	}
}

func (n DIDNamespace) FormatDID(subject string) string {
	return fmt.Sprintf("%s:%s", n.DIDMethod(), subject)
}

type EncodedKey []byte

func (e KeyEncoding) EncodeRaw(data []byte) (EncodedKey, error) {
	switch e {
	case KeyEncoding_KEY_ENCODING_RAW:
		return data, nil
	case KeyEncoding_KEY_ENCODING_HEX:
		return []byte(hex.EncodeToString(data)), nil
	case KeyEncoding_KEY_ENCODING_MULTIBASE:
		return []byte(base58.Encode(data)), nil
	default:
		return nil, nil
	}
}

func (e KeyEncoding) DecodeRaw(data EncodedKey) ([]byte, error) {
	switch e {
	case KeyEncoding_KEY_ENCODING_RAW:
		return data, nil
	case KeyEncoding_KEY_ENCODING_HEX:
		return hex.DecodeString(string(data))
	case KeyEncoding_KEY_ENCODING_MULTIBASE:
		return base58.Decode(string(data))
	default:
		return nil, nil
	}
}

type COSEAlgorithmIdentifier int

func (k KeyAlgorithm) CoseIdentifier() COSEAlgorithmIdentifier {
	switch k {
	case KeyAlgorithm_KEY_ALGORITHM_ES256:
		return COSEAlgorithmIdentifier(-7)
	case KeyAlgorithm_KEY_ALGORITHM_ES384:
		return COSEAlgorithmIdentifier(-35)
	case KeyAlgorithm_KEY_ALGORITHM_ES512:
		return COSEAlgorithmIdentifier(-36)
	case KeyAlgorithm_KEY_ALGORITHM_EDDSA:
		return COSEAlgorithmIdentifier(-8)
	case KeyAlgorithm_KEY_ALGORITHM_ES256K:
		return COSEAlgorithmIdentifier(-10)
	default:
		return COSEAlgorithmIdentifier(0)
	}
}

func (k KeyCurve) ComputePublicKey(data []byte) (*PubKey, error) {
	return nil, ErrUnsupportedKeyCurve
}

func (k *Keyshare) Equals(o crypto.MPCShare) bool {
	opk := o.GetPublicKey()
	if opk != nil && k.PublicKey == nil {
		return false
	}
	return k.GetRole() == o.GetRole()
}

func (k *Keyshare) IsUser() bool {
	return k.Role == 2
}

func (k *Keyshare) IsValidator() bool {
	return k.Role == 1
}

// ComputeOriginTXTRecord generates a fingerprint for a given origin
func ComputeOriginTXTRecord(origin string) string {
	h := sha256.New()
	h.Write([]byte(origin))
	return fmt.Sprintf("v=sonr,o=%s,p=%x", origin, h.Sum(nil))
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
