package highway

import (
	context "context"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/sonr-io/core/ipfs"
	ot "go.buf.build/grpc/go/sonr-io/sonr/object"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	res, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// TODO encrypt based on DID

	var permissions uint32 = 0644 // or whatever you need
	err = ioutil.WriteFile("./temp/file.txt", res, fs.FileMode(permissions))
	if err != nil {
		return nil, err
	}

	file, err := ipfs.GetUnixfsNode("./file.txt")
	if err != nil {
		panic(fmt.Errorf("Could not get File: %s", err))
	}

	cidFile, err := s.ipfs.Unixfs().Add(ctx, file)
	if err != nil {
		panic(fmt.Errorf("Could not add File: %s", err))
	}

	fmt.Printf("Added file to IPFS with CID %s\n", cidFile.String())

	someDirectory, err := ipfs.GetUnixfsNode("./temp/file.txt")
	if err != nil {
		panic(fmt.Errorf("Could not get File: %s", err))
	}

	cidDirectory, err := s.ipfs.Unixfs().Add(ctx, someDirectory)
	if err != nil {
		panic(fmt.Errorf("Could not add Directory: %s", err))
	}

	fmt.Printf("Added directory to IPFS with CID %s\n", cidDirectory.String())

	return nil, ErrMethodUnimplemented
}

// ReadObject reads an object.
func (s *HighwayServer) ReadObject(ctx context.Context, req *ot.MsgReadObject) (*ot.MsgReadObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// UpdateObject updates an object.
func (s *HighwayServer) UpdateObject(ctx context.Context, req *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteObject deletes an object.
func (s *HighwayServer) DeleteObject(ctx context.Context, req *ot.MsgDeleteObject) (*ot.MsgDeleteObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}
