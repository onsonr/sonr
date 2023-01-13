// TODO: go-ipfs-api requires libp2p-core this dependency breaks the build
package ipfs

// import (
// 	"bytes"
// 	"context"
// 	"os"

// 	icore "github.com/ipfs/interface-go-ipfs-core"
// 	"github.com/libp2p/go-libp2p/core/peer"
// 	"github.com/sonr-hq/sonr/pkg/node/config"
// )

// type remoteIpfs struct {
// 	ctx    context.Context
// 	// shell  *shell.Shell
// 	config *config.Config
// }

// func newRemote(ctx context.Context, cnfg *config.Config) (*remoteIpfs, error) {
// 	sh := shell.NewShell(cnfg.RemoteIPFSURL)
// 	return &remoteIpfs{shell: sh, ctx: ctx, config: cnfg}, nil
// }

// func (r *remoteIpfs) Add(data []byte) (string, error) {
// 	bzReader := bytes.NewReader(data)
// 	return r.shell.Add(bzReader)
// }

// func (r *remoteIpfs) Get(hash string) ([]byte, error) {
// 	tmpDir, err := os.MkdirTemp("", hash)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer os.RemoveAll(tmpDir)
// 	err = r.shell.Get(hash, tmpDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Read first file from tmpDir
// 	fs, err := os.ReadDir(tmpDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	f, err := os.ReadFile(fs[0].Name())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return f, nil
// }

// func (r *remoteIpfs) Connect(peers ...string) error {
// 	return r.shell.SwarmConnect(r.ctx, peers...)
// }

// func (r *remoteIpfs) MultiAddrs() string {
// 	return r.config.RemoteIPFSURL
// }

// func (r *remoteIpfs) PeerID() peer.ID {
// 	return ""
// }

// func (r *remoteIpfs) GetDecrypted(cidStr string, pubKey []byte) ([]byte, error) {
// 	return nil, nil
// }

// func (r *remoteIpfs) AddEncrypted(file []byte, pubKey []byte) (string, error) {
// 	return "", nil
// }

// func (r *remoteIpfs) CoreAPI() icore.CoreAPI {
// 	return nil
// }
