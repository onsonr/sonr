package protocol

import "context"

type IPFS interface {
	DagGet(ctx context.Context, ref string, out interface{}) error
	DagPut(ctx context.Context, data interface{}, inputCodec, storeCodec string) (string, error)
	GetData(ctx context.Context, cid string) ([]byte, error)
	PutData(ctx context.Context, data []byte) (string, error)
	PinFile(ctx context.Context, cidstr string) error
	RemoveFile(ctx context.Context, cidstr string) error
}
