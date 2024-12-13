package ipfsapi

import (
	"context"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/onsonr/sonr/crypto/mpc"
)

func addEnclaveIPFS(enclave mpc.Enclave, ipc *rpc.HttpApi) (string, error) {
	jsonEnclave, err := enclave.Marshal()
	if err != nil {
		return "", err
	}
	// Save enclave to IPFS
	cid, err := ipc.Unixfs().Add(context.Background(), files.NewBytesFile(jsonEnclave))
	if err != nil {
		return "", err
	}
	return cid.String(), nil
}
