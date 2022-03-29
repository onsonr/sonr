package util

import (
	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	"github.com/pkg/errors"
	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
)

var (
	logger = golog.Default.Child("internal/common")
)

// OLC returns Open Location code
func OLC(l *types.Location) string {
	return olc.Encode(l.GetLatitude(), l.GetLongitude(), 4)
}

// ** ───────────────────────────────────────────────────────
// ** ─── Payload Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
// PayloadItemFunc is the Map function for PayloadItem
type PayloadItemFunc func(item *types.Payload_Item, index int, total int) error

// IsSingle returns true if the transfer is a single transfer. Error returned
// if No Items present in Payload
func IsSingle(p *types.Payload) (bool, error) {
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
func IsMultiple(p *types.Payload) (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return true, nil
	}
	return false, nil
}

// FileCount returns the number of files in the Payload
func FileCount(p *types.Payload) int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type != types.MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}

// URLCount returns the number of URLs in the Payload
func URLCount(p *types.Payload) int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type == types.MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}
