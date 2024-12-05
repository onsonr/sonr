package ipfs

import "github.com/ipfs/kubo/client/rpc"

type client struct {
	api *rpc.HttpApi
}

func NewClient() (*client, error) {
	api, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return &client{api: api}, nil
}
