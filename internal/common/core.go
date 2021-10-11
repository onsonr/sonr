package common

import (
	"fmt"

	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type NodeImpl interface {
	Profile() (*Profile, error)
	Peer() (*Peer, error)
}

// RPC_SERVER_PORT is the port the RPC service listens on.
// Calculated: (Sister Bday + Dad Bday + Mom Bday) / Mine
const RPC_SERVER_PORT = 26225

// GetProfileFunc returns a function that returns the Profile and error
type GetProfileFunc func() (*Profile, error)

var (
	logger = golog.Child("internal/common")
)

// IsMdnsCompatible returns true if the Connection is MDNS compatible
func (c Connection) IsMdnsCompatible() bool {
	return c == Connection_WIFI || c == Connection_ETHERNET
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

func DefaultLocation() *Location {
	return &Location{
		Latitude:  34.102920,
		Longitude: -118.394190,
	}
}

func (l *Location) OLC() string {
	return olc.Encode(l.GetLatitude(), l.GetLongitude(), 6)
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
				logger.Error("Failed to replace path for Item", err)
				return nil, err
			}
		}
	}
	return p, nil
}

// Buffer returns Peer as a buffer
func (p *Profile) Buffer() ([]byte, error) {
	// Marshal Peer
	data, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}

	// Return Peer as buffer
	return data, nil
}
