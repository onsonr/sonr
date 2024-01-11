package didmethod

import (
	"encoding/hex"
	"fmt"

	"google.golang.org/protobuf/proto"
	"lukechampine.com/blake3"

	identityv1 "github.com/sonrhq/identity/api/v1"
)

// WIDXIdentifier is a type alias for a string
type WIDXIdentifier string

// WIDXMethod is the constant DID method name
const WIDXMethod = "widx"

// NewWIDXIdentifier creates a new WIDXIdentifier given a coin type and address
func NewWIDXIdentifier(coinType uint32, address string) (WIDXIdentifier, error) {
    // Create new WIDXIdentifier struct
    wid:= &identityv1.WalletIdentifier{
        CoinType: coinType,
        Value:   address,
    }
    // Marshal the ProtoBuf into bytes
    widBz, err := proto.Marshal(wid)
    if err != nil {
        return "", err
    }

    // Hex encode the Blake3 hash of the bytes
    widHash := blake3.Sum256(widBz)
    hex := hex.EncodeToString(widHash[:])

    // Format and return the DID
    didStr := fmt.Sprintf("did:%s:%s", WIDXMethod, hex)
    return WIDXIdentifier(didStr), nil
}

