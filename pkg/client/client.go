package client

const (
	// -- Local Blockchain --
	BLOCKCHAIN_REST_LOCAL   = "http://0.0.0.0:26657"
	BLOCKCHAIN_FAUCET_LOCAL = "http://0.0.0.0:4500"
	BLOCKCHAIN_RPC_LOCAL    = "127.0.0.1:9090"

	// -- Dev Blockchain --
	BLOCKCHAIN_FAUCET_DEV = "http://143.198.29.209:8000"
	BLOCKCHAIN_RPC_DEV    = "143.198.29.209:9090"

	// -- Beta Blockchain --
	BLOCKCHAIN_FAUCET_BETA = "http://137.184.190.146:8000"
	BLOCKCHAIN_RPC_BETA    = "137.184.190.146:9090"

	// -- Services --
	IPFS_API_ADDRESS  = "http://164.92.99.233"
	VAULT_API_ADDRESS = "http://164.92.99.233"
)

type ConnEndpointType int

const (
	ConnEndpointType_NONE ConnEndpointType = iota
	ConnEndpointType_LOCAL
	ConnEndpointType_DEV
	ConnEndpointType_BETA
)

type Client struct {
	connType ConnEndpointType
}

func NewClient(t ConnEndpointType) *Client {
	return &Client{
		connType: t,
	}
}

func (c *Client) GetFaucetAddress() string {
	switch c.connType {
	case ConnEndpointType_LOCAL:
		return BLOCKCHAIN_FAUCET_LOCAL
	case ConnEndpointType_DEV:
		return BLOCKCHAIN_FAUCET_DEV
	case ConnEndpointType_BETA:
		return BLOCKCHAIN_FAUCET_BETA
	default:
		return BLOCKCHAIN_FAUCET_LOCAL
	}
}

func (c *Client) GetRPCAddress() string {
	switch c.connType {
	case ConnEndpointType_LOCAL:
		return BLOCKCHAIN_RPC_LOCAL
	case ConnEndpointType_DEV:
		return BLOCKCHAIN_RPC_DEV
	case ConnEndpointType_BETA:
		return BLOCKCHAIN_RPC_BETA
	default:
		return BLOCKCHAIN_RPC_LOCAL
	}
}

func (c *Client) GetAPIAddress() string {
	if c.connType == ConnEndpointType_LOCAL {
		return BLOCKCHAIN_REST_LOCAL
	}
	return ""
}
