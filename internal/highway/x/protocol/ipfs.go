package protocol

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"

	"github.com/ipfs/go-cid"

	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/pkg/store"

	ipfslite "github.com/hsanjuan/ipfs-lite"

	"github.com/sonr-io/sonr/pkg/host"
)

var _ protocol.IPFS = (*IPFSLite)(nil)

// NewIPFSLite creates a new IPFSLite instance with Host Implementation
func NewIPFSLite(ctx context.Context, host host.SonrHost) (*IPFSLite, error) {
	// Create IPFS Peer
	ds := store.NewMemoryStore()
	ipfsLite, err := ipfslite.New(ctx, ds.Batching(), host.Host(), host.Routing(), nil)
	if err != nil {
		return nil, err
	}

	p := &IPFSLite{
		ctx:       ctx,
		node:      host,
		dataStore: ds,
		peer:      ipfsLite,
	}

	p.peer.Bootstrap(ipfslite.DefaultBootstrapPeers())
	return p, nil
}

// IPFSLite is an IPFS Lite implementation of ipfs.Protocol.
type IPFSLite struct {
	ctx       context.Context
	node      host.SonrHost
	dataStore *store.Memory
	peer      *ipfslite.Peer
}

func (i *IPFSLite) PinFile(ctx context.Context, cidstr string) error {
	return errors.New("not supported")
}

func (i *IPFSLite) RemoveFile(ctx context.Context, cidstr string) error {
	// Decode CID from String
	c, err := cid.Decode(cidstr)
	if err != nil {
		return err
	}

	return i.peer.Remove(ctx, c)
}

func (i *IPFSLite) DagGet(ctx context.Context, cidstr string, out interface{}) error {
	// Decode CID from String
	c, err := cid.Decode(cidstr)
	if err != nil {
		return err
	}

	v, _ := i.peer.DAGService.Get(ctx, c)
	if err != nil {
		return err
	}

	out = v
	return nil
}

func (i *IPFSLite) DagPut(ctx context.Context, data interface{}, inputCodec, storeCodec string) (string, error) {
	return "", errors.New("not implemented")
}

func (i *IPFSLite) GetData(ctx context.Context, cidstr string) ([]byte, error) {
	// Decode CID from String
	c, err := cid.Decode(cidstr)
	if err != nil {
		return nil, err
	}

	// Get the file from IPFS
	rsc, err := i.peer.GetFile(ctx, c)
	if err != nil {
		return nil, err
	}

	defer rsc.Close()
	return ioutil.ReadAll(rsc)
}

// PutData puts a file to IPFS and returns the CID.
func (i *IPFSLite) PutData(ctx context.Context, data []byte) (string, error) {
	// Create Reader for Data
	buffer := bytes.NewBuffer(data)

	// Adds file to IPFS
	nd, err := i.peer.AddFile(ctx, buffer, nil)
	if err != nil {
		return "", err
	}

	// Get Back the CID
	c := nd.Cid()
	return c.String(), nil
}
