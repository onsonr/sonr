package resolver

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/x/identity/types"
)

func EncodeCredentialVerificationMethod(cred *crypto.WebauthnCredential, controller string) (*types.VerificationMethod, error) {
	did := fmt.Sprintf("did:key:%s#%s", base64.RawURLEncoding.EncodeToString(cred.PublicKey), base64.RawURLEncoding.EncodeToString(cred.Id))
	pubMb := base64.RawURLEncoding.EncodeToString(cred.PublicKey)
	vmType := crypto.Ed25519KeyType.FormatString()
	meta, err := ConvertAuthenticatorToMap(cred.Authenticator)
	if err != nil {
		return nil, err
	}
	return &types.VerificationMethod{
		Id: did,
		Type: vmType,
		PublicKeyMultibase: pubMb,
		Controller: controller,
		Metadata: types.MapToKeyValueList(meta),
	}, nil
}


func DecodeCredentialVerificationMethod(vm *types.VerificationMethod) (*crypto.WebauthnCredential, error) {
	// Extract the public key from PublicKeyMultibase
	pubKey, err := base64.RawURLEncoding.DecodeString(vm.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %v", err)
	}

	// Extract the credential ID from the verification method ID
	id := strings.Split(vm.Id, "#")[1]
	credID, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return nil, fmt.Errorf("failed to decode credential ID: %v", err)
	}

	// Convert metadata to map and build the WebauthnAuthenticator
	authenticatorMap := types.KeyValueListToMap(vm.Metadata)
	authenticator, err := ConvertAuthenticatorFromMap(authenticatorMap)
	if err != nil {
		return nil, fmt.Errorf("failed to build WebauthnAuthenticator: %v", err)
	}

	// Build the WebauthnCredential
	cred := &crypto.WebauthnCredential{
		Id:           credID,
		PublicKey:    pubKey,
		Authenticator: authenticator,
	}

	return cred, nil
}

func ConvertAuthenticatorToMap(authenticator *crypto.WebauthnAuthenticator) (map[string]string, error) {
	authenticatorMap := make(map[string]string)
	aaguid := base64.StdEncoding.EncodeToString(authenticator.Aaguid)
	authenticatorMap["aaguid"] = aaguid
	signCount := strconv.FormatUint(uint64(authenticator.SignCount), 10)
	authenticatorMap["sign_count"] = signCount
	cloneWarning := strconv.FormatBool(authenticator.CloneWarning)
	authenticatorMap["clone_warning"] = cloneWarning
	return authenticatorMap, nil
}

func ConvertAuthenticatorFromMap(authenticatorMap map[string]string) (*crypto.WebauthnAuthenticator, error) {
	aaguid, err := base64.StdEncoding.DecodeString(authenticatorMap["aaguid"])
	if err != nil {
		return nil, err
	}
	signCount, err := strconv.ParseUint(authenticatorMap["sign_count"], 10, 32)
	if err != nil {
		return nil, err
	}
	cloneWarning, err := strconv.ParseBool(authenticatorMap["clone_warning"])
	if err != nil {
		return nil, err
	}
	authenticator := &crypto.WebauthnAuthenticator{
		Aaguid:      aaguid,
		SignCount:   uint32(signCount),
		CloneWarning: cloneWarning,
	}
	return authenticator, nil
}
