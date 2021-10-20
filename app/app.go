package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/viper"
)

type Sonr struct {
	// Properties
	Ctx   context.Context
	Node  api.NodeImpl
	IsCli bool
}

var instance Sonr

func init() {
	golog.SetStacktraceLimit(2)
}

func Start(req *api.InitializeRequest, isTerminal bool, prefix string) {
	if instance.Node != nil {
		golog.Error("Sonr Instance already active")
		return
	}
	golog.SetPrefix(fmt.Sprintf("[Sonr.%s] ", prefix))

	// Initialize Device
	ctx := context.Background()
	err := req.Parse()
	if err != nil {
		golog.Fatal("Failed to initialize Device", golog.Fields{"error": err})
		os.Exit(1)
	}

	// Create Node
	n, _, err := node.NewNode(ctx, node.SetTerminalMode(isTerminal), node.WithRequest(req))
	if err != nil {
		golog.Fatal("Failed to update Profile for Node", golog.Fields{"error": err})
		os.Exit(1)
	}

	// Set Lib
	instance = Sonr{
		Ctx:   ctx,
		IsCli: isTerminal,
		Node:  n,
	}
	instance.Persist()
}

// Exit handles cleanup on Sonr Node
func Exit(code int) {
	if instance.Node == nil {
		golog.Info("Skipping Exit, instance is nil...")
		return
	}
	golog.Info("Cleaning up on Exit...")
	instance.Node.Close()
	defer instance.Ctx.Done()

	// Check for Full Desktop Node
	if common.IsDesktop() {
		ex, err := os.Executable()
		if err != nil {
			golog.Error("Failed to find Executable, ", err)
			return
		}

		// Delete Executable Path
		exPath := filepath.Dir(ex)
		err = os.RemoveAll(filepath.Join(exPath, "sonr_bitcask"))
		if err != nil {
			golog.Warn("Failed to remove Bitcask, ", err)
		}
		err = viper.SafeWriteConfig()
		if err == nil {
			golog.Info("Wrote new config file to Disk")
		}
		os.Exit(code)
	}
}

// Persist waits for Exit Signal from Terminal
func (sh Sonr) Persist() {
	// Check if CLI Mode
	if common.IsMobile() {
		golog.Info("Skipping Serve, Node is either mobile or non-cli...")
		return
	}

	// Wait for Exit Signal
	golog.Info("- Persisting Node on localhost:26225 -")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
	}()

	// Hold until Exit Signal
	for {
		select {
		case <-c:
			Exit(0)
		case <-sh.Ctx.Done():
			return
		}
	}
}
