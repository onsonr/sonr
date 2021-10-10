package lib

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
	"google.golang.org/protobuf/proto"
)

type sonrLib struct {
	// Properties
	ctx  context.Context
	node *node.Node
}

var (
	instance *sonrLib
)

func init() {
	golog.SetPrefix("[Sonr-Core.lib] ")
	golog.SetStacktraceLimit(2)
	golog.SetFormat("json", "    ")
}

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) {
	ctx := context.Background()
	isDev, req, dOpts, err := parseInitializeRequest(reqBuf)
	if err != nil {
		golog.Fatal("Failed to parse initialize request: %s", err)
		return
	}

	// Initialize Device
	err = device.Init(isDev, dOpts...)
	if err != nil {
		golog.Fatal("Failed to initialize Device", err)
		return
	}

	// Create Node
	n, _, err := node.NewNode(ctx, node.WithRequest(req), node.WithMode(node.Mode_CLIENT))
	if err != nil {
		golog.Fatal("Failed to Create new node", err)
		return
	}

	// Set Lib
	instance = &sonrLib{
		ctx:  ctx,
		node: n,
	}
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
	instance.ctx.Done()
}


// parseInitializeRequest parses the given buffer and returns the proto and fsOptions.
func parseInitializeRequest(buf []byte) (bool, *api.InitializeRequest, []device.FSOption, error) {
	// Unmarshal request
	req := &api.InitializeRequest{}
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
			Path: req.GetDeviceOptions().GetTextileDir(),
			Type: device.Textile,
		})
	}
	return req.GetEnvironment().IsDev(), req, fsOpts, nil
}
