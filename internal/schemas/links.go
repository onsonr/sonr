package schemas

import (
	"errors"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
)

func (as *schemaImpl) CreateLinkPrototype() cidlink.LinkPrototype {
	return cidlink.LinkPrototype{Prefix: cid.Prefix{
		Version:  1,    // Usually '1'.
		Codec:    0x71, // 0x71 means "dag-cbor" -- See the multicodecs table: https://github.com/multiformats/multicodec/
		MhType:   0x13, // 0x20 means "sha2-512" -- See the multicodecs table: https://github.com/multiformats/multicodec/
		MhLength: 64,   // sha2-512 hash has a 64-byte sum.
	}}
}

func (as *schemaImpl) LoadLink(key string) (datamodel.Node, error) {
	/*
		lnk := cidlink.Link{Cid: cid.Cid{
			str: key,
		}}

		// as.linkSys.Load(linking.LinkContext{}, lnk, datamodel.Node)
	*/
	return nil, errors.New("LoadLink unimplemented")
}
