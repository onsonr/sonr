package protocol

import "context"

type IPFS interface {
	DagGet(ref string, out interface{}) error
	DagPut(data interface{}, inputCodec, storeCodec string) (string, error)
	LookUpData(ctx context.Context, cid string, data interface{}) error
	GetData(ctx context.Context, cid string) ([]byte, error)
	PutData(ctx context.Context, data []byte) (string, error)
	PinFile(ctx context.Context, cidstr string) error
	RemoveFile(ctx context.Context, cidstr string) error
}
