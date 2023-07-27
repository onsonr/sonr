package crypto

import (
	"fmt"
	"strings"
)

type DIDPrefix string

func (d DIDPrefix) String() string {
	return string(d)
}

type DIDMethod string

func (d DIDMethod) String() string {
	return string(d)
}
func (d DIDMethod) CoinType() CoinType {
	return CoinTypeFromDidMethod(d.String())
}

type DIDIdentifier string

func (d DIDIdentifier) String() string {
	return string(d)
}

func SplitDID(did string) (DIDPrefix, DIDMethod, DIDIdentifier, error) {
	parts := strings.Split(did, ":")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid did: %s", did)
	}
	return DIDPrefix(parts[0]), DIDMethod(parts[1]), DIDIdentifier(parts[2]), nil
}

func CombineDID(prefix DIDPrefix, method DIDMethod, identifier DIDIdentifier) string {
	return fmt.Sprintf("%s:%s:%s", prefix, method, identifier)
}
