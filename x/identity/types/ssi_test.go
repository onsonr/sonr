package types_test

import (
	"testing"

	"github.com/sonr-hq/sonr/x/identity/types"
)

func TestSSIFormatString(t *testing.T) {
	// Test KeyType
	for _, kt := range []types.KeyType{
		types.KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019,
		types.KeyType_KeyType_ED25519_VERIFICATION_KEY_2018,
		types.KeyType_KeyType_JSON_WEB_KEY_2020,
	} {
		assertExpectedKeyTypeFormatString(t, kt.FormatString(), kt)
	}

	// Test ProofType
	for _, pt := range []types.ProofType{
		types.ProofType_ProofType_ECDSA_SECP256K1_SIGNATURE_2019,
		types.ProofType_ProofType_ED25519_SIGNATURE_2018,
		types.ProofType_ProofType_JSON_WEB_SIGNATURE_2020,
	} {
		assertExpectedProofTypeFormatString(t, pt.FormatString(), pt)
	}

	// Test ServiceType
	for _, st := range []types.ServiceType{
		types.ServiceType_ServiceType_DID_COMM_MESSAGING,
		types.ServiceType_ServiceType_ENCRYPTED_DATA_VAULT,
	} {
		assertExpectedServiceTypeFormatString(t, st.FormatString(), st)
	}
}

func assertExpectedKeyTypeFormatString(t *testing.T, actual string, kt types.KeyType) {
	expected := map[types.KeyType]string{
		types.KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019: "EcdsaSecp256k1VerificationKey2019",
		types.KeyType_KeyType_ED25519_VERIFICATION_KEY_2018:         "Ed25519VerificationKey2018",
		types.KeyType_KeyType_JSON_WEB_KEY_2020:                     "JsonWebKey2020",
		types.KeyType_KeyType_RSA_VERIFICATION_KEY_2018:             "RsaVerificationKey2018",
	}
	if expected[kt] != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func assertExpectedProofTypeFormatString(t *testing.T, actual string, pt types.ProofType) {
	expected := map[types.ProofType]string{
		types.ProofType_ProofType_ECDSA_SECP256K1_SIGNATURE_2019: "EcdsaSecp256k1Signature2019",
		types.ProofType_ProofType_ED25519_SIGNATURE_2018:         "Ed25519Signature2018",
		types.ProofType_ProofType_JSON_WEB_SIGNATURE_2020:        "JsonWebSignature2020",
		types.ProofType_ProofType_RSA_SIGNATURE_2018:             "RsaSignature2018",
	}

	if expected[pt] != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func assertExpectedServiceTypeFormatString(t *testing.T, actual string, st types.ServiceType) {
	expected := map[types.ServiceType]string{
		types.ServiceType_ServiceType_DID_COMM_MESSAGING:   "DidCommMessaging",
		types.ServiceType_ServiceType_ENCRYPTED_DATA_VAULT: "EncryptedDataVault",
	}

	if expected[st] != actual {
		t.Errorf("expected %s got %s", expected, actual)
	}
}
