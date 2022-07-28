package protocol

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"sync"

	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/pkg/store"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"

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
		lock:      sync.Mutex{},
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
	lock      sync.Mutex
}

func (i *IPFSLite) PinFile(ctx context.Context, cidstr string) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	return errors.New("not implemented")
}

func (i *IPFSLite) RemoveFile(ctx context.Context, cidstr string) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	return errors.New("not implemented")
}

func (i *IPFSLite) DagGet(ctx context.Context, ref string, out interface{}) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	return errors.New("not implemented")
}

func (i *IPFSLite) DagPut(ctx context.Context, data interface{}, inputCodec, storeCodec string) (string, error) {
	i.lock.Lock()
	defer i.lock.Unlock()

	return "", errors.New("not implemented")
}

func (i *IPFSLite) GetData(ctx context.Context, cid string) ([]byte, error) {
	i.lock.Lock()
	defer i.lock.Unlock()

	// Decode CID from String
	c, err := DecodeCIDFromString(cid)
	if err != nil {
		return nil, err
	}

	// Get the file from IPFS
	rsc, err := i.peer.GetFile(i.ctx, c)
	if err != nil {
		return nil, err
	}

	defer rsc.Close()
	return ioutil.ReadAll(rsc)
}

// DecodeCIDFromString decodes a CID string to a CID.
func DecodeCIDFromString(s string) (cid.Cid, error) {
	return cid.Decode(s)
}

//// GetObjectSchema returns an object schema.
//func (i *IPFSLite) GetObjectSchema(cid *cid.Cid) (datamodel.Node, error) {
//	// Get the file from IPFS
//	buf, err := i.GetData(cid.String())
//	if err != nil {
//		return nil, err
//	}
//
//	// Create bytes reader
//	serial := bytes.NewReader(buf)
//
//	// Decode IPLD Node
//	np := basicnode.Prototype.Any // Pick a stle for the in-memory data.
//	nb := np.NewBuilder()         // Create a builder.
//	dagjson.Decode(nb, serial)    // Hand the builder to decoding -- decoding will fill it in!
//	n := nb.Build()               // Call 'Build' to get the resulting Node.  (It's immutable!)
//	fmt.Printf("the data decoded was a %s kind\n", n.Kind())
//	fmt.Printf("the length of the node is %d\n", n.Length())
//	return n, nil
//}

// PutData puts a file to IPFS and returns the CID.
func (i *IPFSLite) PutData(ctx context.Context, data []byte) (string, error) {
	i.lock.Lock()
	defer i.lock.Unlock()
	
	// Create Reader for Data
	buffer := bytes.NewBuffer(data)

	// Adds file to IPFS
	nd, err := i.peer.AddFile(i.ctx, buffer, nil)
	if err != nil {
		return "", err
	}

	// Get Back the CID
	c := nd.Cid()
	return c.String(), nil
}

//// PutObjectSchema puts an object schema to IPFS and returns the CID.
//func (i *IPFSLite) PutObjectSchema(doc *st.SchemaDefinition) (*cid.Cid, error) {
//	// Create IPLD Node
//	np := basicnode.Prototype.Any
//	nb := np.NewBuilder()                               // Create a builder.
//	ma, err := nb.BeginMap(int64(len(doc.GetFields()))) // Begin assembling a map.
//	if err != nil {
//		return nil, err
//	}
//
//	// Add each field to the map
//	for _, t := range doc.GetFields() {
//		ma.AssembleKey().AssignString(t.Name)
//		switch t.Field {
//		case st.SchemaKind_STRING:
//			ma.AssembleValue().AssignString("")
//		case st.SchemaKind_INT:
//			ma.AssembleValue().AssignInt(0)
//		case st.SchemaKind_FLOAT:
//			ma.AssembleValue().AssignFloat(0.0)
//		case st.SchemaKind_BOOL:
//			ma.AssembleValue().AssignBool(false)
//		case st.SchemaKind_BYTES:
//			ma.AssembleValue().AssignBytes([]byte{})
//		case st.SchemaKind_LINK:
//			ma.AssembleValue().AssignLink(nil)
//		default:
//			ma.AssembleValue().AssignNull()
//		}
//	}
//
//	// End assembling the map.
//	err = ma.Finish()
//	if err != nil {
//		return nil, err
//	}
//
//	// Build IPLD Node
//	n := nb.Build()
//	buffer := &bytes.Buffer{}
//	err = dagjson.Encode(n, buffer)
//	if err != nil {
//		return nil, err
//	}
//
//	// Adds file to IPFS
//	return i.PutData(buffer.Bytes())
//}
