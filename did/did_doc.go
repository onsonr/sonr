package did

import "github.com/sonr-io/core/did/strings"

const (
	PublicKeyJwk       = "PublicKeyJwk"
	PublicKeyMultibase = "PublicKeyMultibase"
)

var VerificationMethodType = map[string]string{
	"JsonWebKey2020":             PublicKeyJwk,
	"Ed25519VerificationKey2020": PublicKeyMultibase,
}

var ServiceType = []string{
	"LinkedDomains",
	"DIDCommMessaging",
}

func GetVerificationMethodType(vmType string) string {
	return VerificationMethodType[vmType]
}

func IsValidDidServiceType(sType string) bool {
	return strings.Contains(ServiceType, sType)
}
