package header

type Key string

const (
	Authorization Key = "Authorization"

	// User Agent
	Architecture    Key = "Sec-CH-UA-Arch"
	Bitness         Key = "Sec-CH-UA-Bitness"
	FullVersionList Key = "Sec-CH-UA-Full-Version-List"
	Mobile          Key = "Sec-CH-UA-Mobile"
	Model           Key = "Sec-CH-UA-Model"
	Platform        Key = "Sec-CH-UA-Platform"
	PlatformVersion Key = "Sec-CH-UA-Platform-Version"
	UserAgent       Key = "Sec-CH-UA"

	// Sonr Injected
	ChainID     Key = "X-Chain-ID"
	IPFSHost    Key = "X-Host-IPFS"
	SonrAPIURL  Key = "X-Sonr-API"
	SonrgRPCURL Key = "X-Sonr-GRPC"
	SonrRPCURL  Key = "X-Sonr-RPC"
	SonrWSURL   Key = "X-Sonr-WS"
)

func (h Key) String() string {
	return string(h)
}
