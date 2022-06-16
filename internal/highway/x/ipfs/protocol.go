package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sonr-io/sonr/internal/highway/x/store"
	"io/ioutil"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/sonr-io/sonr/pkg/host"
	st "github.com/sonr-io/sonr/x/schema/types"
)

// Protocol leverages the IPFSLite library to provide simple file operations.
type Protocol struct {
	ctx   context.Context
	host  host.SonrHost
	store *store.Store
	*ipfslite.Peer
}

// New creates a new Protocol instance with Host Implementation
func New(ctx context.Context, ds *store.Store, host host.SonrHost) (*Protocol, error) {
	// Create IPFS Peer
	peer, err := ipfslite.New(ctx, ds, host.Host(), host.Routing(), nil)
	if err != nil {
		return nil, err
	}

	p := &Protocol{
		ctx,
		host,
		ds,
		peer,
	}

	p.Bootstrap(ipfslite.DefaultBootstrapPeers())
	return p, nil
}

// DecodeCIDFromString decodes a CID string to a CID.
func DecodeCIDFromString(s string) (cid.Cid, error) {
	return cid.Decode(s)
}

// GetData returns a file from IPFS.
func (i *Protocol) GetData(cid string) ([]byte, error) {
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

// GetObjectSchema returns an object schema from IPFS.
func (i *Protocol) GetObjectSchema(cid *cid.Cid) (datamodel.Node, error) {
	// Get the file from IPFS
	buf, err := i.GetData(cid.String())
	if err != nil {
		return nil, err
	}

	// Create bytes reader
	serial := bytes.NewReader(buf)

	// Decode IPLD Node
	np := basicnode.Prototype.Any // Pick a stle for the in-memory data.
	nb := np.NewBuilder()         // Create a builder.
	dagjson.Decode(nb, serial)    // Hand the builder to decoding -- decoding will fill it in!
	n := nb.Build()               // Call 'Build' to get the resulting Node.  (It's immutable!)
	fmt.Printf("the data decoded was a %s kind\n", n.Kind())
	fmt.Printf("the length of the host is %d\n", n.Length())
	return n, nil
}

// PutData puts a file to IPFS and returns the CID.
func (i *Protocol) PutData(data []byte) (*cid.Cid, error) {
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

// PutObjectSchema puts an object schema to IPFS and returns the CID.
func (i *Protocol) PutObjectSchema(doc *st.SchemaDefinition) (*cid.Cid, error) {
	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder()                               // Create a builder.
	ma, err := nb.BeginMap(int64(len(doc.GetFields()))) // Begin assembling a map.
	if err != nil {
		return nil, err
	}

	// Add each field to the map
	for k, t := range doc.GetFields() {
		ma.AssembleKey().AssignString(k)
		switch t {
		case st.SchemaKind_STRING:
			ma.AssembleValue().AssignString("")
		case st.SchemaKind_INT:
			ma.AssembleValue().AssignInt(0)
		case st.SchemaKind_FLOAT:
			ma.AssembleValue().AssignFloat(0.0)
		case st.SchemaKind_BOOL:
			ma.AssembleValue().AssignBool(false)
		case st.SchemaKind_BYTES:
			ma.AssembleValue().AssignBytes([]byte{})
		case st.SchemaKind_LINK:
			ma.AssembleValue().AssignLink(nil)
		default:
			ma.AssembleValue().AssignNull()
		}
	}

	// End assembling the map.
	err = ma.Finish()
	if err != nil {
		return nil, err
	}

	// Build IPLD Node
	n := nb.Build()
	buffer := &bytes.Buffer{}
	err = dagjson.Encode(n, buffer)
	if err != nil {
		return nil, err
	}

	// Adds file to IPFS
	return i.PutData(buffer.Bytes())
}

// RemoveFile removes a file from IPFS.
func (i *Protocol) RemoveFile(cidstr string) error {
	cid, err := DecodeCIDFromString(cidstr)
	if err != nil {
		return err
	}
	return i.Peer.Remove(i.ctx, cid)
}
