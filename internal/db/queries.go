package db

import "fmt"

// Account queries
func insertAccountQuery(name, address string) string {
	return fmt.Sprintf(`INSERT INTO accounts (name, address) VALUES (%s, %s)`, name, address)
}

// Asset queries
func insertAssetQuery(name, symbol string, decimals int, chainID int64) string {
	return fmt.Sprintf(
		`INSERT INTO assets (name, symbol, decimals, chain_id) VALUES (%s, %s, %d, %d)`,
		name,
		symbol,
		decimals,
		chainID,
	)
}

// Chain queries
func insertChainQuery(name string, networkID string) string {
	return fmt.Sprintf(`INSERT INTO chains (name, network_id) VALUES (%s, %d)`, name, networkID)
}

// Credential queries
func insertCredentialQuery(
	handle, controller, attestationType, origin string,
	credentialID, publicKey []byte,
	transport string,
	signCount uint32,
	userPresent, userVerified, backupEligible, backupState, cloneWarning bool,
) string {
	return fmt.Sprintf(`INSERT INTO credentials (
		handle, controller, attestation_type, origin, 
		credential_id, public_key, transport, sign_count, 
		user_present, user_verified, backup_eligible, 
		backup_state, clone_warning
	) VALUES (%s, %s, %s, %s, %s, %s, %s, %d, %t, %t, %t, %t, %t)`,
		handle, controller, attestationType, origin,
		credentialID, publicKey, transport, signCount,
		userPresent, userVerified, backupEligible,
		backupState, cloneWarning)
}

// Profile queries
func insertProfileQuery(
	id, subject, controller, originURI, publicMetadata, privateMetadata string,
) string {
	return fmt.Sprintf(`INSERT INTO profiles (
		id, subject, controller, origin_uri, 
		public_metadata, private_metadata
	) VALUES (%s, %s, %s, %s, %s, %s)`,
		id, subject, controller, originURI,
		publicMetadata, privateMetadata)
}

// Property queries
func insertPropertyQuery(profileID, key, accumulator, propertyKey string) string {
	return fmt.Sprintf(`INSERT INTO properties (
		profile_id, key, accumulator, property_key
	) VALUES (%s, %s, %s, %s)`,
		profileID, key, accumulator, propertyKey)
}

// Permission queries
func insertPermissionQuery(serviceID, grants, scopes string) string {
	return fmt.Sprintf(
		`INSERT INTO permissions (service_id, grants, scopes) VALUES (%s, %s, %s)`,
		serviceID,
		grants,
		scopes,
	)
}

// GetPermission query
func getPermissionQuery(serviceID string) string {
	return fmt.Sprintf(`SELECT grants, scopes FROM permissions WHERE service_id = %s`, serviceID)
}
