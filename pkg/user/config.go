package user

import (
	"fmt"
	"path/filepath"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/internal/network"
	md "github.com/sonr-io/core/pkg/models"
)

type UserConfig struct {
	connectivity  md.Connectivity
	isDesktop     bool
	hasPrivateKey bool
	privateKey    crypto.PrivKey
	fileSystem    *FileSystem
	user          *md.User
	settings      []*md.User_Settings
}

// ^ Initialize User ^ //
func InitUserConfig(req *md.ConnectionRequest, cb md.NodeCallback) *UserConfig {
	// Initialize
	var sonrPath string
	var hasPrivateKey bool

	// Check for Client Type
	if req.Device.GetIsDesktop() {
		// Init Path, Check for Path
		sonrPath = filepath.Join(req.Directories.Home, K_SONR_CLIENT_DIR)
		if err := EnsureDir(sonrPath, 0755); err != nil {
			return nil
		}
	} else {
		// Set Path to Documents for Mobile
		sonrPath = req.Directories.Documents
	}

	// Set File System
	fs := &FileSystem{
		IsDesktop: req.Device.GetIsDesktop(),
		Downloads: req.Directories.Downloads,
		Main:      sonrPath,
		Temporary: req.Directories.Temporary,
		Call:      cb,
	}

	// Get Private Key
	privKey, err := fs.getPrivateKey()
	if err != nil {
		hasPrivateKey = false
	} else {
		hasPrivateKey = true
	}

	// Build Config
	return &UserConfig{
		connectivity:  req.Connectivity,
		isDesktop:     req.Device.IsDesktop,
		fileSystem:    fs,
		privateKey:    privKey,
		hasPrivateKey: hasPrivateKey,
	}
}

// ^ Return Networking Options ^ //
func (u *UserConfig) HostOptions() libp2p.Option {
	// IP Address
	ip4 := network.IPv4()
	ip6 := network.IPv6()

	// Set Defaults
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip6/%s/tcp/0", ip6),
		),
		libp2p.NATPortMap(),
		libp2p.EnableAutoRelay(),
	}

	// Check for Private Key
	if u.hasPrivateKey {
		opts = append(opts, libp2p.Identity(u.privateKey))
	}
	return libp2p.ChainOptions(opts...)
}
