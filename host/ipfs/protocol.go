package ipfs

import (
	"context"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/sonr-io/core/host"
)

// IPFSProtocol leverages the IPFSLite library to provide simple file operations.
type IPFSProtocol struct {
	ctx       context.Context
	node      host.HostImpl
	dataStore datastore.Batching
	*ipfslite.Peer
}

// New creates a new IPFSProtocol instance with Host Implementation
func New(ctx context.Context, host host.HostImpl) (*IPFSProtocol, error) {
	// TODO - Create a better batching.Batching data store opposed to just in-memory
	ds := ipfslite.NewInMemoryDatastore()
	ipfsLite, err := ipfslite.New(ctx, ds, host.Host(), host.Routing(), nil)
	if err != nil {
		return nil, err
	}

	p := &IPFSProtocol{
		ctx:       ctx,
		node:      host,
		dataStore: ds,
		Peer:      ipfsLite,
	}

	p.Bootstrap(ipfslite.DefaultBootstrapPeers())
	return p, nil
}

// DecodeCIDFromString decodes a CID string to a CID.
func (i *IPFSProtocol) DecodeCIDFromString(s string) (cid.Cid, error) {
	return cid.Decode(s)
}

// GetFile returns a file from IPFS.
func (i *IPFSProtocol) GetFile(cid cid.Cid) ([]byte, error) {
	rsc, err := i.GetFile(cid)
	if err != nil {
		return nil, err
	}
	return rsc, nil
}

// PutFile puts a file to IPFS and returns the CID.
func (i *IPFSProtocol) PutFile(data []byte) (cid.Cid, error) {
	cid, err := i.PutFile(data)
	if err != nil {
		return cid, err
	}
	return cid, nil
}

// RemoveFile removes a file from IPFS.
func (i *IPFSProtocol) RemoveFile(cid cid.Cid) error {
	return i.RemoveFile(cid)
}
