package protocol

import "context"

type IPFS interface {
	GetData(ctx context.Context, cid string) ([]byte, error)
	PutData(ctx context.Context, data []byte) (string, error)
	PinFile(ctx context.Context, cidstr string) error
	RemoveFile(ctx context.Context, cidstr string) error
}
