package remote

import (
	"bytes"
	"context"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p/core/peer"
)

var (
	DefaultAPIAddr = "/ip4/198.199.78.62/tcp/9094"
)

type RemoteIPFS struct {
	ctx   context.Context
	shell *shell.Shell
}

func NewApi(ctx context.Context) (*RemoteIPFS, error) {
	sh := shell.NewShell(DefaultAPIAddr)
	return &RemoteIPFS{shell: sh, ctx: ctx}, nil
}

func (r *RemoteIPFS) Add(data []byte) (string, error) {
	bzReader := bytes.NewReader(data)
	return r.shell.Add(bzReader)
}

func (r *RemoteIPFS) Get(hash string) ([]byte, error) {
	tmpDir, err := os.MkdirTemp("", hash)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)
	err = r.shell.Get(hash, tmpDir)
	if err != nil {
		return nil, err
	}
	// Read first file from tmpDir
	fs, err := os.ReadDir(tmpDir)
	if err != nil {
		return nil, err
	}
	f, err := os.ReadFile(fs[0].Name())
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *RemoteIPFS) Connect(peers ...string) error {
	return r.shell.SwarmConnect(r.ctx, peers...)
}

func (r *RemoteIPFS) MultiAddr() string {
	return DefaultAPIAddr
}

func (r *RemoteIPFS) PeerID() peer.ID {
	return ""
}

func (r *RemoteIPFS) GetDecrypted(cidStr string, pubKey []byte) ([]byte, error) {
	return nil, nil
}

func (r *RemoteIPFS) AddEncrypted(file []byte, pubKey []byte) (string, error) {
	return "", nil
}

func (r *RemoteIPFS) CoreAPI() icore.CoreAPI {
	return nil
}
