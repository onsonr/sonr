package types

// DefaultAssets returns the default asset infos: BTC, ETH, SNR, and USDC
func DefaultAssets() []*AssetInfo {
	return []*AssetInfo{
		{
			Name:      "Bitcoin",
			Symbol:    "BTC",
			Hrp:       "bc",
			Index:     0,
			AssetType: AssetType_ASSET_TYPE_NATIVE,
			IconUrl:   "https://cdn.sonr.land/BTC.svg",
		},
		{
			Name:      "Ethereum",
			Symbol:    "ETH",
			Hrp:       "eth",
			Index:     64,
			AssetType: AssetType_ASSET_TYPE_NATIVE,
			IconUrl:   "https://cdn.sonr.land/ETH.svg",
		},
		{
			Name:      "Sonr",
			Symbol:    "SNR",
			Hrp:       "idx",
			Index:     703,
			AssetType: AssetType_ASSET_TYPE_NATIVE,
			IconUrl:   "https://cdn.sonr.land/SNR.svg",
		},
	}
}

// DefaultChains returns the default chain infos: Bitcoin, Ethereum, and Sonr.
func DefaultChains() []*ChainInfo {
	return []*ChainInfo{}
}

// DefaultKeyInfos returns the default key infos: secp256k1, ed25519, keccak256, and bls12377.
func DefaultKeyInfos() []*KeyInfo {
	return []*KeyInfo{
		//
		// Identity Key Info
		//
		// Sonr Controller Key Info
		{
			Role:      KeyRole_KEY_ROLE_INVOCATION,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256K,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
		},
		{
			Role:      KeyRole_KEY_ROLE_ASSERTION,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_BLS12377,
			Encoding:  KeyEncoding_KEY_ENCODING_MULTIBASE,
		},

		//
		// Blockchain Key Info
		//
		// Ethereum Key Info
		{
			Role:      KeyRole_KEY_ROLE_DELEGATION,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_KECCAK256,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
		},
		// Bitcoin Key Info
		{
			Role:      KeyRole_KEY_ROLE_DELEGATION,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256K,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
		},

		//
		// Authentication Key Info
		//
		// Browser based WebAuthn
		{
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
		},
		// FIDO U2F
		{
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256K,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
		},
		// Cross-Platform Passkeys
		{
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_EDDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
		},
	}
}

func DefaultOpenIDConfig() *OpenIDConfig {
	return &OpenIDConfig{
		Issuer:                 "https://sonr.id",
		AuthorizationEndpoint:  "https://api.sonr.id/auth",
		TokenEndpoint:          "https://api.sonr.id/token",
		UserinfoEndpoint:       "https://api.sonr.id/userinfo",
		ScopesSupported:        []string{"openid", "profile", "email", "web3", "sonr"},
		ResponseTypesSupported: []string{"code"},
		ResponseModesSupported: []string{"query", "form_post"},
		GrantTypesSupported:    []string{"authorization_code", "refresh_token"},
		AcrValuesSupported:     []string{"passkey"},
		SubjectTypesSupported:  []string{"public"},
	}
}
