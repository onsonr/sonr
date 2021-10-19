package api

import (
	"fmt"
	"os"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/internal/wallet"
)

// DefaultInitializeRequest returns the default initialize request
func DefaultInitializeRequest() *InitializeRequest {
	return &InitializeRequest{
		Profile:  common.NewDefaultProfile(),
		Location: common.DefaultLocation(),
	}
}

// FSOpts returns a list of FS Options
func (ir *InitializeRequest) FSOpts() []fs.Option {
	return []fs.Option{
		fs.WithHomePath(ir.homeDir()),
		fs.WithSupportPath(ir.supportDir()),
		fs.WithTempPath(ir.tempDir()),
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

// SetEnvVars sets the environment variables
func (ir *InitializeRequest) Parse() error {
	// Set Environment Variables
	vars := ir.GetVariables()
	count := len(vars)

	// Iterate over Variables
	if count > 0 {
		for k, v := range vars {
			os.Setenv(k, v)
		}

		golog.Info("Added Enviornment Variable(s)", golog.Fields{
			"Total": count,
		})
	}

	// Set Device ID
	did := ir.GetDeviceOptions().GetId()
	if did != "" {
		logger.Info("Device ID Passed: " + did)
		common.SetDeviceID(did)
	}

	// Start File System
	if err := fs.Start(ir.FSOpts()...); err != nil {
		return errors.Wrap(err, "Failed to Start File System")
	}

	// Open Keychain
	if err := wallet.Open(); err != nil {
		return errors.Wrap(err, "Failed to Open Keychain")
	}
	return nil
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
		} else if fs.IsFile(item.GetPath()) {
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
