package common

import (
	"fmt"
	"runtime"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// GetProfileFunc returns a function that returns the Profile and error
type GetProfileFunc func() (*Profile, error)

// Buffer returns Peer as a buffer
func (p *Peer) Buffer() ([]byte, error) {
	// Marshal Peer
	data, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}

	// Return Peer as buffer
	return data, nil
}

// Arch returns Peer Device GOARCH
func (p *Peer) Arch() string {
	return p.GetDevice().GetArch()
}

// Libp2pID returns the PeerID based on PublicKey from Profile
func (p *Peer) Libp2pID() (peer.ID, error) {
	// Check if PublicKey is empty
	if len(p.GetPublicKey()) == 0 {
		return "", errors.New("Peer Public Key is not set.")
	}

	pubKey, err := crypto.UnmarshalPublicKey(p.GetPublicKey())
	if err != nil {
		return "", err
	}

	// Return Peer ID
	id, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	return id, nil
}

// PubKey returns the Public Key from the Peer
func (p *Peer) PubKey() (crypto.PubKey, error) {
	// Check if PublicKey is empty
	if len(p.GetPublicKey()) == 0 {
		return nil, errors.New("Peer Public Key is not set.")
	}

	// Unmarshal Public Key
	pubKey, err := crypto.UnmarshalPublicKey(p.GetPublicKey())
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to Unmarshal Public Key: %s", p.GetSName()), err)
		return nil, err
	}
	return pubKey, nil
}

// OS returns Peer Device GOOS
func (p *Peer) OS() string {
	return p.GetDevice().GetOs()
}

// Add adds a new Profile to the List and
// updates LastModified time.
func (p *ProfileList) Add(profile *Profile) {
	p.Profiles = append(p.Profiles, profile)
	p.LastModified = time.Now().Unix()
}

// Count returns the number of Profiles in the List
func (p *ProfileList) Count() int {
	return len(p.Profiles)
}

// IndexAt returns profile at index
func (p *ProfileList) IndexAt(i int) *Profile {
	return p.Profiles[i]
}

// DefaultProfileOption is a type for Profile Options
type DefaultProfileOption func(profileOpts)

// profileOpts contains the options for generating a new Profile
type profileOpts struct {
	refProfile *Profile
	sname      string
	firstName  string
	lastName   string
	picture    []byte
	bio        string
	socials    []*Social
}

// defaultProfileOpts returns the default profile options
func defaultProfileOpts() profileOpts {
	return profileOpts{
		sname:     fmt.Sprintf("a%s", runtime.GOOS),
		firstName: "Anonymous",
		lastName:  runtime.GOOS,
		picture:   make([]byte, 0),
		socials:   make([]*Social, 0),
	}
}

// NewDefaultProfile creates a new default Profile
func NewDefaultProfile(options ...DefaultProfileOption) *Profile {
	// Set default options
	opts := defaultProfileOpts()
	for _, option := range options {
		option(opts)
	}

	// Create Profile Build Func
	buildProfile := func(opts profileOpts) *Profile {
		return &Profile{
			SName:        opts.sname,
			FirstName:    opts.firstName,
			LastName:     opts.lastName,
			Picture:      opts.picture,
			Bio:          opts.bio,
			Socials:      opts.socials,
			LastModified: time.Now().Unix(),
		}
	}

	// Check if refProfile is set and build
	if opts.refProfile != nil {
		if !checkProfile(opts.refProfile) {
			return buildProfile(opts)
		}
		return opts.refProfile
	}
	return buildProfile(opts)
}

// WithCheckerProfile sets the checker profile
func WithCheckerProfile(profile *Profile) DefaultProfileOption {
	return func(opts profileOpts) {
		opts.refProfile = profile
	}
}

// checkProfile checks if the Profile is valid
func checkProfile(p *Profile) bool {
	if p == nil {
		return false
	}
	if len(p.SName) == 0 || len(p.FirstName) == 0 {
		return false
	}
	return true
}
