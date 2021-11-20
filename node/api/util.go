package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/device"
)

// IsIncoming returns true if the event is incoming
func (e *CompleteEvent) IsIncoming() bool {
	return e.GetDirection() == common.Direction_INCOMING
}

// IsOutgoing returns true if the event is outgoing
func (e *CompleteEvent) IsOutgoing() bool {
	return e.GetDirection() == common.Direction_OUTGOING
}

// Recent returns the profile of CompleteEvent by Direction
func (e *CompleteEvent) Recent() *common.Profile {
	if e.Direction == common.Direction_INCOMING {
		return e.GetFrom().GetProfile()
	}
	return e.GetTo().GetProfile()
}

// Title returns the title of the event
func (d *CompleteEvent) Title() string {
	return fmt.Sprintf("[Transfer-Complete] from %s at %v", d.GetFrom().GetProfile().GetSName(), time.Now())
}

// Message returns the message of the event
func (d *CompleteEvent) Message() string {
	paths := ""
	for _, v := range d.GetPayload().GetItems() {
		paths += fmt.Sprintf("\n-\t%s", v.GetFile().GetPath())
	}
	return fmt.Sprintf("Size: %v \n Paths: %s", d.GetPayload().GetSize(), paths)
}

// Title returns the title of the event
func (d *DecisionEvent) Title() string {
	return fmt.Sprintf("[Transfer-Decision] from %s", d.GetFrom().GetProfile().GetSName())
}

// Message returns the message of the event
func (d *DecisionEvent) Message() string {
	return fmt.Sprintf("Result: %v", d.GetDecision())
}

// Title returns the title of the event
func (d *InviteEvent) Title() string {
	fname := d.GetFrom().GetProfile().GetFirstName()
	lname := d.GetFrom().GetProfile().GetLastName()
	sname := d.GetFrom().GetProfile().GetSName()
	platform := strings.ToUpper(d.GetFrom().GetDevice().GetOs())
	return fmt.Sprintf("[Transfer-Invite] from %s %s (%s) on (%s)", fname, lname, sname, platform)
}

// Message returns the message of the event
func (d *InviteEvent) Message() string {
	fcount := d.GetPayload().FileCount()
	ucount := d.GetPayload().URLCount()
	tcount := fcount + ucount
	countStr := fmt.Sprintf("%d, (Files) %d, (Urls) %d", tcount, fcount, ucount)
	mimes := ""
	for _, v := range d.GetPayload().GetItems() {
		mimes += fmt.Sprintf("%s, ", v.GetMime().GetValue())
	}
	return fmt.Sprintf("Count: %s \nMimes: %s \nSize: %v", countStr, mimes, d.GetPayload().GetSize())
}

// DefaultInitializeRequest returns the default initialize request
func DefaultInitializeRequest() *InitializeRequest {
	return &InitializeRequest{
		Profile:  common.NewDefaultProfile(),
		Location: DefaultLocation(),
	}
}

// DefaultLocation is the default location if the location is not found
func DefaultLocation() *common.Location {
	return &common.Location{
		Latitude:  34.102920,
		Longitude: -118.394190,
	}
}

// Options returns a list of FS Options
func (ir *InitializeRequest) Options() []device.Option {
	return []device.Option{
		device.WithHomePath(ir.homeDir()),
		device.WithSupportPath(ir.supportDir()),
		device.WithTempPath(ir.tempDir()),
		device.SetDeviceID(ir.GetDeviceOptions().GetId()),
	}
}

// homeDir returns provided String Home Path
func (ir *InitializeRequest) homeDir() string {
	return ir.GetDeviceOptions().GetHomeDir()
}

// supportDir returns provided String Support Path
func (ir *InitializeRequest) supportDir() string {
	return ir.GetDeviceOptions().GetSupportDir()
}

// tempDir returns provided String Temporary Path
func (ir *InitializeRequest) tempDir() string {
	return ir.GetDeviceOptions().GetTempDir()
}

// IsDev returns true if the node is running in development mode.
func (ir *InitializeRequest) IsDev() bool {
	return ir.GetEnvironment().IsDev()
}

// Count returns the number of items in the payload
func (sr *ShareRequest) Count() int {
	return len(sr.GetItems())
}

// ToPayload converts the response to a payload
func (sr *ShareRequest) ToPayload(owner *common.Profile) (*common.Payload, error) {
	// Initialize
	fileCount := 0
	urlCount := 0
	size := int64(0)
	items := make([]*common.Payload_Item, 0)

	// Iterate over Paths
	for i, item := range sr.GetItems() {
		// Check if path is URL
		if common.IsUrl(item.GetPath()) {
			// Increase URL Count
			urlCount++

			// Add URL to Payload
			urlItem, err := common.NewUrlItem(item.GetPath())
			if err != nil {
				msg := fmt.Sprintf("Failed to create URLItem at Index: %v, with Path: %s", i, item.GetPath())
				logger.Error(msg, err)
				return nil, errors.Wrap(err, msg)
			}

			// Add URL to Payload
			items = append(items, urlItem)
		} else if device.IsFile(item.GetPath()) {
			// Increase File Count
			fileCount++

			// Create Payload Item
			fileItem, err := common.NewFileItem(item.GetPath(), item.GetThumbnail())
			if err != nil {
				msg := fmt.Sprintf("Failed to create FileItem at Index: %v with Path: %s", i, item.GetPath())
				logger.Error(msg, err)
				return nil, errors.Wrap(err, msg)
			}

			// Add Payload Item to Payload
			items = append(items, fileItem)
			size += fileItem.GetSize()
		}
	}

	// Log Payload Details
	logger.Debug(fmt.Sprintf("Created payload with %v Files and %v URLs. Total size: %v", fileCount, urlCount, size))

	// Create Payload
	payload := &common.Payload{
		Items: items,
		Size:  size,
		Owner: owner,
	}
	return payload, nil
}

// NewInitialzeResponse returns a new initialize response
func NewInitialzeResponse(gpf common.GetProfileFunc, success bool) *InitializeResponse {
	resp := &InitializeResponse{Success: success}
	if !success || gpf == nil {
		return resp
	}
	p, err := gpf()
	if err != nil {
		logger.Errorf("%s - Failed to get profile", err)
		return resp
	}
	resp.Profile = p
	return resp
}
