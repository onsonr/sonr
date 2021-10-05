package common

import (
	"fmt"
	"strings"
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/logger"
)

// ** ───────────────────────────────────────────────────────
// ** ─── General ───────────────────────────────────────────
// ** ───────────────────────────────────────────────────────
// OLC_SCOPE is the default OLC Scope for Distance Calculation
const OLC_SCOPE = 4

// Fetch olc code from lat/lng at Scope Level 6
func (l *Location) OLC() string {
	return olc.Encode(float64(l.GetLatitude()), float64(l.GetLongitude()), OLC_SCOPE)
}

// Checks if Enviornment is Development
func (e Environment) IsDev() bool {
	return e == Environment_DEVELOPMENT
}

// Checks if Enviornment is Development
func (e Environment) IsProd() bool {
	return e == Environment_PRODUCTION
}

// WrapErrors wraps errors list into a single error
func WrapErrors(msg string, errs []error) error {
	// Check if errors are empty
	if len(errs) == 0 {
		return nil
	}

	// Iterate over errors
	err := errors.New(msg)
	for _, e := range errs {
		if e != nil {
			err = errors.Wrap(e, e.Error())
			continue
		}
	}
	return err
}

// ** ───────────────────────────────────────────────────────
// ** ─── Peer Management ───────────────────────────────────
// ** ───────────────────────────────────────────────────────

// PeerInfo is a struct for Peer Information containing Device and Crypto
type PeerInfo struct {
	OperatingSystem string        // Device Operating System
	Architecture    string        // Device Architecture
	HostName        string        // Device Host Name
	SName           string        // Peer SName
	StoreEntryKey   string        // Peer SName in Store Entry Key Format
	PeerID          peer.ID       // Peer ID
	Peer            *Peer         // Peer Data Object
	PublicKey       crypto.PubKey // Peer Public Key
}

// Info returns PeerInfo from Peer
func (p *Peer) Info() (*PeerInfo, error) {
	// Get Public Key
	pubKey, err := p.PubKey()
	if err != nil {
		logger.Error("Failed to get Public Key", err)
		return nil, err
	}

	// Get peer ID from public key
	id, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to get peer ID from Public Key: %s", p.GetSName()), err)
		return nil, err
	}

	// Return Peer Info
	return &PeerInfo{
		OperatingSystem: p.GetDevice().GetOs(),
		Architecture:    p.GetDevice().GetArch(),
		HostName:        p.GetDevice().GetHostName(),
		PeerID:          id,
		PublicKey:       pubKey,
		SName:           p.GetSName(),
		StoreEntryKey:   strings.ToLower(p.GetSName()),
		Peer:            p,
	}, nil
}

// PeerID returns the PeerID based on PublicKey from Profile
func (p *Peer) PeerID() (peer.ID, error) {
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
		return nil, logger.Error(fmt.Sprintf("Failed to Unmarshal Public Key: %s", p.GetSName()), err)
	}
	return pubKey, nil
}

// SnrPubKey returns the Public Key from the Peer as SnrPubKey
func (p *Peer) SnrPubKey() (*keychain.SnrPubKey, error) {
	// Get Public Key
	pub, err := p.PubKey()
	if err != nil {
		return nil, logger.Error("Failed to get Public Key", err)
	}

	// Return SnrPubKey
	return keychain.NewSnrPubKey(pub), nil
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

// ** ───────────────────────────────────────────────────────
// ** ─── Payload Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
// PayloadItemFunc is the Map function for PayloadItem
type PayloadItemFunc func(item *Payload_Item, index int, total int) error

// NewPayload creates a new Payload Object
func NewPayload(owner *Profile, paths []string) (*Payload, error) {
	// Initialize
	fileCount := 0
	urlCount := 0
	size := int64(0)
	items := make([]*Payload_Item, 0)
	errs := make([]error, 0)

	// Iterate over Paths
	for i, path := range paths {
		// Check if path is URL
		if IsUrl(path) {
			// Increase URL Count
			urlCount++

			// Add URL to Payload
			item, err := NewUrlItem(path)
			if err != nil {
				msg := fmt.Sprintf("Failed to create URLItem at Index: %v, with Path: %s", i, path)
				logger.Error(msg, err)
				errs = append(errs, errors.Wrap(err, msg))
				continue
			}

			// Add URL to Payload
			items = append(items, item)
			continue
		} else if IsFile(path) {
			// Increase File Count
			fileCount++

			// Create Payload Item
			item, err := NewFileItem(path)
			if err != nil {
				msg := fmt.Sprintf("Failed to create FileItem at Index: %v with Path: %s", i, path)
				logger.Error(msg, err)
				errs = append(errs, errors.Wrap(err, msg))
				continue
			}

			// Add Payload Item to Payload
			items = append(items, item)
			size += item.GetSize()
			continue
		} else {
			err := fmt.Errorf("Invalid Path provided, value is neither File or URL. Path: %s", path)
			logger.Error(err.Error(), err)
			errs = append(errs, err)
			continue
		}
	}

	// Log Payload Details
	logger.Info(fmt.Sprintf("Created payload with %v Files and %v URLs. Total size: %v", fileCount, urlCount, size))

	// Create Payload
	payload := &Payload{
		Items: items,
		Size:  size,
		Owner: owner,
	}

	// Check if there are any errors
	if len(errs) > 0 {
		err := WrapErrors(fmt.Sprintf("⚠️ Payload created with %v Errors: \n", len(errs)), errs)
		logger.Error(err.Error(), err)
		return payload, err
	}
	return payload, nil
}

// IsSingle returns true if the transfer is a single transfer. Error returned
// if No Items present in Payload
func (p *Payload) IsSingle() (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return false, nil
	}
	return true, nil
}

// IsMultiple returns true if the transfer is a multiple transfer. Error returned
// if No Items present in Payload
func (p *Payload) IsMultiple() (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return true, nil
	}
	return false, nil
}

// MapItems performs method chaining on the Items in the Payload
func (p *Payload) MapItems(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if err := fn(item, i, count); err != nil {
			return err
		}
	}
	return nil
}

// MapItems performs method chaining on the Items in the Payload
func (p *Payload) MapItemsWithIndex(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if err := fn(item, i, count); err != nil {
			return err
		}
	}
	return nil
}

// MapFileItems performs method chaining on ONLY the FileItems in the Payload
func (p *Payload) MapFileItems(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if item.GetFile() != nil {
			if err := fn(item, i, count); err != nil {
				return err
			}
		}
	}
	return nil
}

// MapUrlItems performs method chaining on ONLY the UrlItems in the Payload
func (p *Payload) MapUrlItems(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if item.GetUrl() != nil {
			if err := fn(item, i, count); err != nil {
				return err
			}
		}
	}
	return nil
}

// ReplaceItemsDir iterates over the items in the payload and replaces the
// directory of the item with the new directory.
func (p *Payload) ReplaceItemsDir(dir string) (*Payload, error) {
	// Create new Payload
	for _, item := range p.GetItems() {
		if item.GetFile() != nil {
			err := item.GetFile().ReplaceDir(dir)
			if err != nil {
				return nil, logger.Error("Failed to replace path for Item", err)
			}
		}
	}
	return p, nil
}
