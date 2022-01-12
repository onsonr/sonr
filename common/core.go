package common

import (
	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

var (
	logger = golog.Default.Child("internal/common")
)

// NewDefaultLocation returns the Sonr HQ as default location
func NewDefaultLocation() *Location {
	return &Location{
		Latitude:  float64(40.673010),
		Longitude: float64(-73.994450),
		Placemark: &Location_Placemark{
			Name:                  "Sonr HQ",
			Street:                "94 9th St.",
			IsoCountryCode:        "US",
			Country:               "United States",
			AdministrativeArea:    "New York",
			SubAdministrativeArea: "Brooklyn",
			Locality:              "Brooklyn",
			SubLocality:           "Gowanus",
			PostalCode:            "11215",
		},
	}
}

// IsMdnsCompatible returns true if the Connection is MDNS compatible
func (c Connection) IsMdnsCompatible() bool {
	return c == Connection_CONNECTION_WIFI || c == Connection_CONNECTION_ETHERNET
}

// IsDev Checks if Enviornment is Development
func (e Environment) IsDev() bool {
	return e == Environment_ENVIRONMENT_DEVELOPMENT
}

// IsProd Checks if Enviornment is Development
func (e Environment) IsProd() bool {
	return e == Environment_ENVIRONMENT_PRODUCTION
}

// OLC returns Open Location code
func (l *Location) OLC() string {
	return olc.Encode(l.GetLatitude(), l.GetLongitude(), 4)
}

// ** ───────────────────────────────────────────────────────
// ** ─── Payload Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
// PayloadItemFunc is the Map function for PayloadItem
type PayloadItemFunc func(item *Payload_Item, index int, total int) error

// IsFile returns true if the Item is a File
func (p *Payload) IsFile() bool {
	isFile := false
	for _, item := range p.GetItems() {
		isFile = item.GetMime().IsFile()
	}
	return isFile
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

// FileCount returns the number of files in the Payload
func (p *Payload) FileCount() int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type != MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
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

// URLCount returns the number of URLs in the Payload
func (p *Payload) URLCount() int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type == MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
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
