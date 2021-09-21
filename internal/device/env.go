package device

import (
	"fmt"
	"os"
)

// HDNS client key for Namebase.io
var HANDSHAKE_KEY string

// HDNS secret key for Namebase.io
var HANDSHAKE_SECRET string

// Rapid API location key
var IP_LOCATION_KEY string

// Rapid API key
var RAPID_API_KEY string

// Textile Hub API key
var TEXTILE_HUB_KEY string

// Textile Hub API secret
var TEXTILE_HUB_SECRET string

func initEnv() error {
	// Set environment variables
	HANDSHAKE_KEY = os.Getenv("HANDSHAKE_KEY")
	HANDSHAKE_SECRET = os.Getenv("HANDSHAKE_SECRET")
	IP_LOCATION_KEY = os.Getenv("IP_LOCATION_KEY")
	RAPID_API_KEY = os.Getenv("RAPID_API_KEY")
	TEXTILE_HUB_KEY = os.Getenv("TEXTILE_HUB_KEY")
	TEXTILE_HUB_SECRET = os.Getenv("TEXTILE_HUB_SECRET")

	// Check for missing environment variables
	if HANDSHAKE_KEY == "" {
		return envVarError("HANDSHAKE_KEY")
	}
	if HANDSHAKE_SECRET == "" {
		return envVarError("HANDSHAKE_SECRET")
	}
	if IP_LOCATION_KEY == "" {
		return envVarError("IP_LOCATION_KEY")
	}
	if RAPID_API_KEY == "" {
		return envVarError("RAPID_API_KEY")
	}
	if TEXTILE_HUB_KEY == "" {
		return envVarError("TEXTILE_HUB_KEY")
	}
	if TEXTILE_HUB_SECRET == "" {
		return envVarError("TEXTILE_HUB_SECRET")
	}
	return nil
}

func envVarError(name string) error {
	return fmt.Errorf("Missing Env Variable for: %s", name)
}
