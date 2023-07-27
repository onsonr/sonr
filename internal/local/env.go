package local

import (
	"fmt"
	"os"
	"time"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 Default Values                                 ||
// ! ||--------------------------------------------------------------------------------||

const (
	kDefaultChainID               = "sonr-localnet-1"
	kDefaultValidatorAddress      = "0x0000000000"
	kDefaultAccountIceFireEnabled = false

	kDefaultHighwayHostPort       = ":8080"
	kDefaultHighwayRequestTimeout = 15

	kDefaultIceFireHost   = "localhost:6001"
	kDefaultJWTSigningKey = "sercrethatmaycontainch@r$32chars"

	kDefaultNodeAPIHost  = "0.0.0.0:1317"
	kDefaultNodeGrpcHost = "0.0.0.0:9090"
	kDefaultNodeRpcHost  = "0.0.0.0:26657"
	kDefaultTLSCertPath  = ""
	kDefaultTLSKeyPath   = ""
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                              Environment Variables                             ||
// ! ||--------------------------------------------------------------------------------||

// ChainID returns the chain ID from the environment variable SONR_CHAIN_ID. (default: sonr-localnet-1)
func ChainID() string {
	if env := os.Getenv("SONR_CHAIN_ID"); env != "" {
		return env
	}
	return kDefaultChainID
}

// HighwayHostPort returns the node highway host from the environment variable SONR_NODE_HIGHWAY_HOST. (default: :8080)
func HighwayHostPort() string {
	if env := os.Getenv("SONR_NODE_HIGHWAY_HOST"); env != "" {
		return env
	}
	return kDefaultHighwayHostPort
}

// HighwayRequestTimeout returns the node highway request timeout from the environment variable SONR_NODE_HIGHWAY_REQUEST_TIMEOUT. (default: 8sec)
func HighwayRequestTimeout() time.Duration {
	if env := os.Getenv("SONR_NODE_HIGHWAY_REQUEST_TIMEOUT"); env != "" {
		if timeout, err := time.ParseDuration(env); err == nil {
			return timeout
		}
	}
	return time.Second * kDefaultHighwayRequestTimeout
}

// IceFireHost returns the IceFire host from the environment variable SONR_ICEFIRE_HOST. (default: 0.0.0.0:6001)
func IceFireHost() string {
	if env := os.Getenv("SONR_ICEFIRE_HOST"); env != "" {
		return env
	}
	return kDefaultIceFireHost
}

// JWTSigningKey returns the JWT signing key from the environment variable SONR_JWT_SIGNING_KEY. (default: sercrethatmaycontainch@r$32chars)
func JWTSigningKey() []byte {
	if env := os.Getenv("SONR_JWT_SIGNING_KEY"); env != "" {
		return []byte(env)
	}
	return []byte(kDefaultJWTSigningKey)
}

// NodeAPIHost returns the node API host from the environment variable SONR_NODE_API_HOST. (default: 0.0.0.0:1317)
func NodeAPIHost() string {
	if env := os.Getenv("SONR_NODE_API_HOST"); env != "" {
		return env
	}
	return kDefaultNodeAPIHost
}

// NodeGrpcHost returns the node gRPC host from the environment variable SONR_NODE_GRPC_HOST. (default: 0.0.0.0:9090)
func NodeGrpcHost() string {
	if env := os.Getenv("SONR_NODE_GRPC_HOST"); env != "" {
		return env
	}
	return kDefaultNodeGrpcHost
}

// NodeRpcHost returns the node RPC host from the environment variable SONR_NODE_RPC_HOST. (default: 0.0.0.0:26657)
func NodeRpcHost() string {
	if env := os.Getenv("SONR_NODE_RPC_HOST"); env != "" {
		return env
	}
	return kDefaultNodeRpcHost
}

// Environment returns the environment from the environment variable SONR_ENVIRONMENT. (default: local)
func Environment() string {
	if env := os.Getenv("SONR_ENVIRONMENT"); env != "" {
		return env
	}
	return "local"
}

// IsProduction returns true if the environment is production.
func IsProduction() bool {
	return Environment() == "production"
}

// IsIceFireEnabled returns true if the account icefire is enabled.
func IsIceFireEnabled() bool {
	if env := os.Getenv("SONR_ACCOUNT_ICEFIRE_ENABLED"); env != "" {
		return env == "true"
	}
	return kDefaultAccountIceFireEnabled
}

// PublicDomain returns the public domain from the environment variable SONR_PUBLIC_DOMAIN. (default: localhost)
func PublicDomain() string {
	if env := os.Getenv("SONR_PUBLIC_DOMAIN"); env != "" {
		return env
	}
	return "localhost"
}

// PublicDomainURLs returns the preconfigured list of domains to use for autotls configuration in production.
func PublicDomainURLs() []string {
	return []string{
		PublicDomain(),
		fmt.Sprintf("api.%s", PublicDomain()),
		fmt.Sprintf("rpc.%s", PublicDomain()),
		fmt.Sprintf("grpc.%s", PublicDomain()),
	}
}

// ValidatorAddress returns the validator address from the environment variable SONR_VALIDATOR_ADDRESS. (default: 0x0000000000)
func ValidatorAddress() string {
	if env := os.Getenv("SONR_VALIDATOR_ADDRESS"); env != "" {
		return env
	}
	return kDefaultValidatorAddress
}
