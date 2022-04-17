package ipfs

import (
	"bytes"
	"context"
	"io/ioutil"

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
func (i *IPFSProtocol) GetFile(cid string) ([]byte, error) {
	validCid, err := i.DecodeCIDFromString(cid)
	if err != nil {
		return nil, err
	}
	rsc, err := i.Peer.GetFile(i.ctx, validCid)
	if err != nil {
		return nil, err
	}
	defer rsc.Close()
	return ioutil.ReadAll(rsc)
}

// PutFile puts a file to IPFS and returns the CID.
func (i *IPFSProtocol) PutFile(data []byte) (*cid.Cid, error) {
	// Create Reader for Data
	buffer := bytes.NewBuffer(data)

	// Adds file to IPFS
	nd, err := i.Peer.AddFile(i.ctx, buffer, nil)
	if err != nil {
		return nil, err
	}

	// Get Back the CID
	c := nd.Cid()
	return &c, nil
}

// RemoveFile removes a file from IPFS.
func (i *IPFSProtocol) RemoveFile(cid string) error {
	validCid, err := i.DecodeCIDFromString(cid)
	if err != nil {
		return err
	}
	return i.Peer.Remove(i.ctx, validCid)
}
