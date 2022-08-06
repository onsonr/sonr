package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"
	"github.com/sonr-io/sonr/pkg/host"
	"github.com/sonr-io/sonr/pkg/protocol"
)

// IPFSProtocol leverages the IPFSLite library to provide simple file operations.
type IPFSProtocol struct {
	ctx       context.Context
	node      host.SonrHost
	dataStore *MemoryStore
	*ipfslite.Peer
}

// New creates a new IPFSProtocol instance with Host Implementation
func New(ctx context.Context, host host.SonrHost) (protocol.IPFS, error) {
	// Create IPFS Peer
	ds := NewMemoryStore()
	ipfsLite, err := ipfslite.New(ctx, ds.Batching(), host.Host(), host.Routing(), nil)
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
func DecodeCIDFromString(s string) (cid.Cid, error) {
	return cid.Decode(s)
}

// GetData returns a file from IPFS.
func (i *IPFSProtocol) GetData(ctx context.Context, cid string) ([]byte, error) {
	// Decode CID from String
	c, err := DecodeCIDFromString(cid)
	if err != nil {
		return nil, err
	}

	// Get the file from IPFS
	rsc, err := i.Peer.GetFile(i.ctx, c)
	if err != nil {
		return nil, err
	}

	defer rsc.Close()
	return ioutil.ReadAll(rsc)
}

// PutData puts a file to IPFS and returns the CID.
func (i *IPFSProtocol) PutData(ctx context.Context, data []byte) (string, error) {
	// Create Reader for Data
	buffer := bytes.NewBuffer(data)

	// Adds file to IPFS
	nd, err := i.Peer.AddFile(i.ctx, buffer, nil)
	if err != nil {
		return "", err
	}

	// Get Back the CID
	c := nd.Cid()
	return c.String(), nil
}

func (i *IPFSProtocol) PinFile(ctx context.Context, cidstr string) error {
	return fmt.Errorf("Not supported")
}

func (i *IPFSProtocol) DagGet(ctx context.Context, ref string, out interface{}) error {
	_, cid, err := cid.CidFromBytes([]byte(ref))
	if err != nil {
		return err
	}

	n, err := i.Peer.DAGService.Get(ctx, cid)

	if err != nil {
		return err
	}

	out = n.RawData()

	return nil
}

func (i *IPFSProtocol) DagPut(ctx context.Context, data interface{}, inputCodec, storeCodec string) (string, error) {
	return "", nil
}

// RemoveFile removes a file from IPFS.
func (i *IPFSProtocol) RemoveFile(ctx context.Context, cidstr string) error {
	cid, err := DecodeCIDFromString(cidstr)
	if err != nil {
		return err
	}
	return i.Peer.Remove(i.ctx, cid)
}
