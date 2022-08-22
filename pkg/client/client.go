package client

import st "github.com/sonr-io/sonr/x/schema/types"

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
	IPFS_ADDRESS      = "https://ipfs.sonr.ws"
	IPFS_API_ADDRESS  = "https://api.ipfs.sonr.ws"
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
	connType          ConnEndpointType
	

	whatIsStore map[string]*st.WhatIs
	schemaStore map[string]*st.SchemaDefinition
}

func NewClient(t ConnEndpointType) *Client {
	return &Client{
		connType:    t,
		whatIsStore: make(map[string]*st.WhatIs),
		schemaStore: make(map[string]*st.SchemaDefinition),
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

func (c *Client) GetIPFSAddress() string {
	return IPFS_ADDRESS
}

func (c *Client) GetIPFSApiAddress() string {
	return IPFS_API_ADDRESS
}
