package didmethod

import (
	"encoding/hex"
	"fmt"

	"google.golang.org/protobuf/proto"
	"lukechampine.com/blake3"

	identityv1 "github.com/sonrhq/identity/api/v1"
)

// UIDXIdentifier is a type alias for a string
type UIDXIdentifier string

// UIDXMethod is the constant DID method name
const UIDXMethod = "uidx"

// NewUIDXIdentifier creates a new WIDXIdentifier given a coin type and address
func NewUIDXIdentifier(idType identityv1.UserIdentifierType, value string) (UIDXIdentifier, error) {
    // Create new WIDXIdentifier struct
    wid:= &identityv1.UserIdentifier{
        IdentifierType: idType,
        Value:   value,
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
    didStr := fmt.Sprintf("did:%s:%s", UIDXMethod, hex)
    return UIDXIdentifier(didStr), nil
}
