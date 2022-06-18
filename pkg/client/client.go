package client

const (
	// HTTP Faucet Address - Public
	SONR_HTTP_FAUCET_PUBLIC = "http://143.198.29.209:8000"

	// HTTP Faucet Address - Local
	SONR_HTTP_FAUCET_LOCAL = "http://0.0.0.0:4500"

	// RPC Address for public node
	SONR_RPC_ADDR_PUBLIC = "143.198.29.209:9090"

	// RPC Address for local node
	SONR_RPC_ADDR_LOCAL = "127.0.0.1:9090"

	// Sonr REST API Address - LOCAL ONLY
	SONR_REST_API_ADDR_LOCAL = "http://0.0.0.0:26657"
)

type Client struct {
	IsLocal bool
}

func NewClient(isLocal bool) *Client {
	return &Client{
		IsLocal: isLocal,
	}
}

func (c *Client) GetFaucetAddress() string {
	if c.IsLocal {
		return SONR_HTTP_FAUCET_LOCAL
	}
	return SONR_HTTP_FAUCET_PUBLIC
}

func (c *Client) GetRPCAddress() string {
	if c.IsLocal {
		return SONR_RPC_ADDR_LOCAL
	}
	return SONR_RPC_ADDR_PUBLIC
}

func (c *Client) GetAPIAddress() string {
	if c.IsLocal {
		return SONR_REST_API_ADDR_LOCAL
	}
	return ""
}
