package orm

const SCHEMA_VERSION = 1

func AccountSchema() string {
	return "++, id, name, address, publicKey, chainCode, index, controller, createdAt"
}

func AssetSchema() string {
	return "++, id, name, symbol, decimals, chainCode, createdAt"
}

func ChainSchema() string {
	return "++, id, name, networkId, chainCode, createdAt"
}

func CredentialSchema() string {
	return "++, id, subject, controller, attestationType, origin, label, deviceId, credentialId, publicKey, transport, signCount, userPresent, userVerified, backupEligible, backupState, cloneWarning, createdAt, updatedAt"
}

func DIDSchema() string {
	return "++, id, role, algorithm, encoding, curve, key_type, raw, jwk"
}

func JwkSchema() string {
	return "++, kty, crv, x, y, n, e"
}

func GrantSchema() string {
	return "++, subject, controller, origin, token, scopes, createdAt, updatedAt"
}

func KeyshareSchema() string {
	return "++, id, data, role, createdAt, lastRefreshed"
}

func ProfileSchema() string {
	return "++, id, subject, controller, originUri, publicMetadata, privateMetadata, createdAt, updatedAt"
}
