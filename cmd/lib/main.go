package lib

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/logger"
	"google.golang.org/protobuf/proto"
)

type SonrLib struct {
	// Properties
	ctx  context.Context
	node *node.Node
}

var sonrLib *SonrLib

// Start starts the host, node, and rpc service.
func Start(reqBytes []byte) []byte {
	// Unmarshal request
	isDev, req, fsOpts, err := parseInitializeRequest(reqBytes)
	if err != nil {
		logger.Panic("Failed to Parse Initialize Request", err)
	}

	// Initialize Device
	ctx := context.Background()
	err = device.Init(isDev, fsOpts...)
	if err != nil {
		logger.Panic("Failed to initialize Device", err)
	}

	// Create Node
	n, resp, err := node.NewNode(ctx, node.WithRequest(req), node.WithClient(), node.WithEnvMap(req.GetVariables()))
	if err != nil {
		logger.Panic("Failed to update Profile for Node", err)
	}

	// Set Lib
	sonrLib = &SonrLib{
		ctx:  ctx,
		node: n,
	}

	// Marshal Response
	buf, err := proto.Marshal(resp)
	if err != nil {
		logger.Warn("Failed to Marshal InitializeResponse", err)
	}
	return buf
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	// if started {
	// 	// state.GetState().Pause()
	// }
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	// if started {
	// 	// state.GetState().Resume()
	// }
}

// Stop closes the host, node, and rpc service.
func Stop() {
	// if started {
	// 	client.host.Close()
	// 	client.ctx.Done()
	// }
}

// parseInitializeRequest parses the given buffer and returns the proto and fsOptions.
func parseInitializeRequest(buf []byte) (bool, *node.InitializeRequest, []device.FSOption, error) {
	// Unmarshal request
	req := &node.InitializeRequest{}
	err := proto.Unmarshal(buf, req)
	if err != nil {
		return false, nil, nil, err
	}

	// Check FSOptions and Get Device Paths
	fsOpts := make([]device.FSOption, 0)
	if req.GetDeviceOptions() != nil {
		// Set Device ID
		err = device.SetDeviceID(req.GetDeviceOptions().GetId())
		if err != nil {
			return req.GetEnvironment().IsDev(), nil, nil, err
		}

		// Set Temporary Path
		fsOpts = append(fsOpts, device.FSOption{
			Path: req.GetDeviceOptions().GetCacheDir(),
			Type: device.Temporary,
		}, device.FSOption{
			Path: req.GetDeviceOptions().GetDownloadsDir(),
			Type: device.Downloads,
		}, device.FSOption{
			Path: req.GetDeviceOptions().GetDocumentsDir(),
			Type: device.Documents,
		}, device.FSOption{
			Path: req.GetDeviceOptions().GetSupportDir(),
			Type: device.Support,
		}, device.FSOption{
			Path: req.GetDeviceOptions().GetDatabaseDir(),
			Type: device.Database,
		}, device.FSOption{
			Path: req.GetDeviceOptions().GetMailboxDir(),
			Type: device.Mailbox,
		})
	}
	return req.GetEnvironment().IsDev(), req, fsOpts, nil
}
