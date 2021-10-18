package api

import (
	"fmt"
	"os"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/config"
	"google.golang.org/protobuf/encoding/protojson"
)

// DefaultInitializeRequest returns the default initialize request
func DefaultInitializeRequest() *InitializeRequest {
	return &InitializeRequest{
		Profile:  common.NewDefaultProfile(),
		Location: common.DefaultLocation(),
	}
}

// IsDev returns true if the node is running in development mode.
func (ir *InitializeRequest) IsDev() bool {
	return ir.GetEnvironment().IsDev()
}

// fsOpts returns the device.FSOptions
func (ir *InitializeRequest) fsOpts() []device.FSOption {
	fsOpts := make([]device.FSOption, 0)
	if ir.GetDeviceOptions() != nil {
		// Set Temporary Path
		fsOpts = append(fsOpts, device.FSOption{
			Path: ir.GetDeviceOptions().GetCacheDir(),
			Type: device.Temporary,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetDownloadsDir(),
			Type: device.Downloads,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetDocumentsDir(),
			Type: device.Documents,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetSupportDir(),
			Type: device.Support,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetDatabaseDir(),
			Type: device.Database,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetTextileDir(),
			Type: device.Textile,
		})
	}
	return fsOpts
}

// MarshalJSON marshals the request to JSON
func (ir *InitializeRequest) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(ir)
}

// UnmarshalJSON unmarshals the request from JSON
func (ir *InitializeRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, ir)
}

// ParseOpts parses the device options and returns the device.FSOptions
func (ir *InitializeRequest) ParseOpts() []device.FSOption {
	logger.Info("Parsing Initialize Request...")

	// Check DeviceID
	ir.SetEnvVars()
	ir.SetDeviceID()
	return ir.fsOpts()
}

// SetEnvVars sets the environment variables
func (ir *InitializeRequest) SetEnvVars() {
	vars := ir.GetVariables()
	count := len(vars)

	// Set Env Variables
	if count > 0 {
		for k, v := range vars {
			os.Setenv(k, v)
		}

		golog.Info("Added Enviornment Variable(s)", golog.Fields{
			"Total": count,
		})
	} else {
		golog.Warn("No Enviornment Variable(s) passed")
	}
}

// SetDeviceID sets the device id
func (ir *InitializeRequest) SetDeviceID() {
	did := ir.GetDeviceOptions().GetId()
	if did != "" {
		logger.Info("Device ID Passed: " + did)
		device.SetDeviceID(did)
	} else {
		golog.Warn("No Device ID Passed")
	}
}

// IsDelete returns true if the request is a delete request
func (er *EditRequest) IsDelete() bool {
	return er.GetType() == EditRequest_DELETE
}

func (sr *SupplyRequest) Count() int {
	return len(sr.GetItems())
}

func (sr *SupplyRequest) ToPayload(owner *common.Profile) (*common.Payload, error) {
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
		} else if config.IsFile(item.GetPath()) {
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
	logger.Info(fmt.Sprintf("Created payload with %v Files and %v URLs. Total size: %v", fileCount, urlCount, size))

	// Create Payload
	payload := &common.Payload{
		Items: items,
		Size:  size,
		Owner: owner,
	}
	return payload, nil
}
