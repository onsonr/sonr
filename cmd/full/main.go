package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kataras/golog"
	"github.com/pterm/pterm"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
	"google.golang.org/protobuf/encoding/protojson"
)

type Sonr struct {
	// Properties
	ctx  context.Context
	node api.NodeImpl
}

var (
	snr *Sonr
)

func init() {
	golog.SetPrefix("[Sonr-Core.highway] ")
	golog.SetStacktraceLimit(2)
	pterm.SetDefaultOutput(golog.Default.Printer)
}

func main() {
	// Parse InitializeRequest
	req, err := Parse()
	if err != nil {
		golog.Warn("Failed to Parse Initialize Request, Using Default Request", golog.Fields{"error": err.Error()})
		req = api.DefaultInitializeRequest()
	}

	// Initialize Device
	deviceSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).WithShowTimer(true).Start("Initializing Device...")
	ctx := context.Background()
	err = device.Init()
	deviceSpinner.Stop()
	if err != nil {
		golog.Fatal("Failed to initialize Device", golog.Fields{"error": err})
	}

	// Create Node
	nodeSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).WithShowTimer(true).Start("Starting Full Node...")
	n, _, err := node.NewNode(ctx, node.WithTerminal(), node.WithRequest(req))
	nodeSpinner.Stop()
	if err != nil {
		golog.Fatal("Failed to update Profile for Node", golog.Fields{"error": err})
	}

	// Set Lib
	snr = &Sonr{
		ctx:  ctx,
		node: n,
	}
	snr.Serve()
}

func AppHeader(s *Sonr) *pterm.TextPrinter {
	p, err := s.node.Peer()
	if err != nil {
		golog.Error("Failed to get Peer", golog.Fields{"error": err})
		s.Exit(1)
		return nil
	}
	header := fmt.Sprintf("Node Available: %v", p.PeerID)
	return pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		header)
}

func (sh *Sonr) Serve() {
	AppHeader(sh)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		snr.Exit(0)
	}()
	for {
		select {
		case <-sh.ctx.Done():
			golog.Info("Context Done")
			return
		}
	}
}

func (sh *Sonr) Exit(code int) {
	golog.Info("Cleaning up on Exit...")
	defer sh.ctx.Done()
	ex, err := os.Executable()
	if err != nil {
		golog.Error("Failed to find Executable, ", err)
	}

	exPath := filepath.Dir(ex)
	err = os.RemoveAll(filepath.Join(exPath, "sonr_bitcask"))
	if err != nil {
		golog.Error("Failed to remove Bitcask, ", err)
	}
	os.Exit(code)
}

// Parse parses the given request and returns Request
func Parse() (*api.InitializeRequest, error) {
	// Get DefaultInitializeRequest
	defReq := api.DefaultInitializeRequest()
	buf, err := defReq.MarshalJSON()
	if err != nil {
		return defReq, err
	}

	// Parse Flag
	req := &api.InitializeRequest{}
	reqPtr := flag.String("req", string(buf), "InitializeRequest JSON String")
	flag.Parse()

	// Unmarshal Request
	err = protojson.Unmarshal([]byte(*reqPtr), req)
	if err != nil {
		return defReq, err
	}
	req.SetEnvVars()
	return req, nil
}
