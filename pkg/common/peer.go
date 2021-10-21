package common

import (
	"bytes"
	"fmt"
	"image/png"
	"math/rand"
	"strings"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/o1egl/govatar"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/wallet"
	"google.golang.org/protobuf/proto"
)

// ** ───────────────────────────────────────────────────────
// ** ─── Peer Management ───────────────────────────────────
// ** ───────────────────────────────────────────────────────

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

	// Fetch public key from peer data
	pubKey, err := p.SnrPubKey()
	if err != nil {
		return "", err
	}

	// Return Peer ID
	id, err := pubKey.PeerID()
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

// SnrPubKey returns the Public Key from the Peer as SnrPubKey
func (p *Peer) SnrPubKey() (*wallet.SnrPubKey, error) {
	// Get Public Key
	pub, err := p.PubKey()
	if err != nil {
		logger.Error("Failed to get Public Key", err)
		return nil, err
	}

	// Return SnrPubKey
	return wallet.NewSnrPubKey(pub), nil
}

// ** ───────────────────────────────────────────────────────
// ** ─── Profile Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
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
	f := faker.FirstName()
	l := faker.LastName()
	return profileOpts{
		sname:     strings.ToLower(f[0:1] + l),
		firstName: "Anonymous",
		lastName:  Platform(),
		picture:   make([]byte, 0),
		bio:       faker.Dessert(),
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

// WithPicture adds a random Profile Picture
func WithPicture() DefaultProfileOption {
	return func(opts profileOpts) {
		opts.picture = genAvatar()
	}
}

// WithSocials adds random Social Media profiles
func WithSocials() DefaultProfileOption {
	return func(opts profileOpts) {
		socials := make([]*Social, 0)
		for i := 0; i < 5; i++ {
			socials = append(socials, genSocial())
		}
		opts.socials = socials
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

// genAvatar generates a random avatar returns empty byte list if error
func genAvatar() []byte {
	// Generate a random avatar
	img, err := govatar.Generate(govatar.MALE)
	if err != nil {
		return make([]byte, 0)
	}

	// Write Img to byte list
	buff := new(bytes.Buffer)
	err = png.Encode(buff, img)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}
	return buff.Bytes()
}

// genSocial generates a random social
func genSocial() *Social {
	mediaIdx := rand.Intn(len(Social_Media_value)-1) + 1
	media := Social_Media(mediaIdx)
	username := faker.Username()
	return &Social{
		Valid:    true,
		Media:    Social_Media(mediaIdx),
		Username: username,
		Url:      fmt.Sprintf("https://%s.com/%s", media.String(), username),
		Picture:  genAvatar(),
	}
}
