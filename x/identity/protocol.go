package identity

import (
	"context"
	"strings"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/api"
	"github.com/sonr-io/core/pkg/device"
	"github.com/sonr-io/core/pkg/wallet"
	"github.com/sonr-io/core/x/common"
	"google.golang.org/protobuf/proto"
)

type IdentityProtocol struct {
	ctx   context.Context
	host  *host.SNRHost
	node  api.NodeImpl
	mode  api.StubMode
	store *bitcask.Bitcask
}

// New creates a new IdentityProtocol
func New(ctx context.Context, host *host.SNRHost, node api.NodeImpl, options ...Option) (*IdentityProtocol, error) {
	// Open the my.db data file in your current directory.
	path, err := device.Database.GenPath("sonr_bitcask")
	if err != nil {
		logger.Errorf("Failed to generate path for bitcask: %s", err)
		return nil, err
	}

	// It will be created if it doesn't exist.
	db, err := bitcask.Open(path, bitcask.WithAutoRecovery(true))
	if err != nil {
		logger.Errorf("%s - Failed to open Database", err)
		return nil, err
	}

	// Create Exchange Protocol
	protocol := &IdentityProtocol{
		ctx:   ctx,
		host:  host,
		node:  node,
		store: db,
	}

	// Set Default Options
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	// Apply Options
	if err := opts.Apply(protocol); err != nil {
		logger.Errorf("%s - Failed to apply options", err)
		return nil, err
	}
	logger.Debug("✅  IdentityProtocol is Activated \n")
	return protocol, nil
}

// AddRecent stores the profile for recents in desk and returns list of recent profiles
func (p *IdentityProtocol) AddRecent(profile *common.Profile) error {
	// Put in Bucket
	if p.store.Has(recentsKey()) {
		// Get profile list buffer
		plist, err := p.addToProfileList(profile)
		if err != nil {
			return err
		}

		// Marshal profile
		val, err := proto.Marshal(plist)
		if err != nil {
			logger.Errorf("%s - Failed to Marshal ProfileList")
			return err
		}

		err = p.store.Put(recentsKey(), val)
		if err != nil {
			return err
		}
		return nil
	}
	profileList := common.ProfileList{
		Key: string(recentsKey()),
	}
	profileList.Add(profile)
	val, err := proto.Marshal(&profileList)
	if err != nil {
		return err
	}
	err = p.store.Put(recentsKey(), val)
	if err != nil {
		return err
	}
	return nil
}

// GetRecents returns the list of recent profiles
func (n *IdentityProtocol) GetRecents() (*common.ProfileList, error) {
	// Check for Key
	if n.store.Has(recentsKey()) {
		rbuf, err := n.store.Get(recentsKey())
		if err != nil {
			logger.Errorf("%s - Failed to Get Recents from store")
			return nil, err
		}

		// Unmarshal profile list
		profileList := common.ProfileList{}
		err = proto.Unmarshal(rbuf, &profileList)
		if err != nil {
			logger.Errorf("%s - Failed to Unmarshal ProfileList")
			return nil, err
		}
		return &profileList, nil
	}
	return &common.ProfileList{}, nil
}

// Peer method returns the peer of the node
func (p *IdentityProtocol) Peer() (*common.Peer, error) {
	// Get Profile
	profile, err := p.Profile()
	if err != nil {
		logger.Warn("Failed to get profile from Memory store, using DefaultProfile.", err)
	}

	// Get Public Key
	pubKey, err := wallet.Sonr.GetSnrPubKey(wallet.Account)
	if err != nil {
		logger.Errorf("%s - Failed to get Public Key", err)
		return nil, err
	}

	// Marshal Public Key
	pubBuf, err := pubKey.Buffer()
	if err != nil {
		logger.Errorf("%s - Failed to marshal public key", err)
		return nil, err
	}

	stat, err := device.Stat()
	if err != nil {
		logger.Errorf("%s - Failed to get device stat", err)
		return nil, err
	}
	// Return Peer
	return &common.Peer{
		SName:        strings.ToLower(profile.GetSName()),
		Status:       common.Peer_ONLINE,
		Profile:      profile,
		PublicKey:    pubBuf,
		PeerID:       p.host.ID().String(),
		LastModified: time.Now().Unix(),
		Device: &common.Peer_Device{
			HostName: stat["hostName"],
			Os:       stat["os"],
			Id:       stat["id"],
			Arch:     stat["arch"],
		},
	}, nil
}

// Profile returns the profile for the user from diskDB
func (n *IdentityProtocol) Profile() (*common.Profile, error) {
	if n.store.Has(PROFILE_KEY) {
		pbuf, err := n.store.Get(PROFILE_KEY)
		if err != nil {
			return common.NewDefaultProfile(), err
		}

		profile := common.Profile{}
		err = proto.Unmarshal(pbuf, &profile)
		if err != nil {
			return nil, err
		}
		return &profile, nil
	}
	return common.NewDefaultProfile(), nil
}

// SetProfile stores the profile for the user in diskDB
func (n *IdentityProtocol) SetProfile(profile *common.Profile) error {
	// Check if profile is nil
	if profile == nil {
		return ErrMissingParam
	}
	pbuf, err := profile.Buffer()
	if err != nil {
		return err
	}
	err = n.store.Put(PROFILE_KEY, pbuf)
	if err != nil {
		return err
	}
	return nil
}
