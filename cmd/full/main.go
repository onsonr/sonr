package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/kataras/golog"
	"github.com/pterm/pterm"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
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
	golog.SetPrefix("[Sonr-Core.full] ")
	golog.SetStacktraceLimit(2)
	pterm.SetDefaultOutput(golog.Default.Printer)
}

func main() {
	// Parse InitializeRequest
	req, err := handleInitRequest()
	if err != nil {
		golog.Warn("Failed to Parse Initialize Request, Using Default Request", golog.Fields{"error": err.Error()})
		req = api.DefaultInitializeRequest()
	}

	// Initialize Device
	ctx := context.Background()
	err = req.Parse()
	if err != nil {
		golog.Fatal("Failed to initialize Device", golog.Fields{"error": err})
		os.Exit(1)
	}

	// Create Node
	n, _, err := node.NewNode(ctx, node.WithTerminal(), node.WithRequest(req))
	if err != nil {
		golog.Fatal("Failed to update Profile for Node", golog.Fields{"error": err})
		os.Exit(1)
	}

	// Set Lib
	snr = &Sonr{
		ctx:  ctx,
		node: n,
	}
	snr.Serve()
}

// AppHeader prints Node Info onto Terminal
func AppHeader(s *Sonr) {
	p, err := s.node.Peer()
	if err != nil {
		golog.Error("Failed to get Peer", golog.Fields{"error": err})
		s.Exit(1)

	}

	//
	pterm.DefaultSection.Println(fmt.Sprintf("Sonr Node Online: %s", p.PeerID))
	pterm.Info.Println(fmt.Sprintf("SName: %s \nOS: %s \nArch: %s", p.GetSName(), p.OS(), p.Arch()))
}

// Serve waits for Exit Signal from Terminal
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

// Exit handles cleanup on Sonr A[[]]
func (sh *Sonr) Exit(code int) {
	golog.Info("Cleaning up on Exit...")
	sh.node.Close()
	defer sh.ctx.Done()
	ex, err := os.Executable()
	if err != nil {
		golog.Error("Failed to find Executable, ", err)
		return
	}

	exPath := filepath.Dir(ex)
	golog.Info(filepath.Join(exPath, "sonr_bitcask"))
	err = os.RemoveAll(filepath.Join(exPath, "sonr_bitcask"))
	if err != nil {
		golog.Error("Failed to remove Bitcask, ", err)
	}
	os.Exit(code)
}

// handleInitRequest parses the given request and returns Request
func handleInitRequest() (*api.InitializeRequest, error) {
	// Parse flag
	latPtr := flag.Float64("lat", 34.102920, "Latitude for InitializeRequest")
	lngPtr := flag.Float64("lng", -118.394190, "Longitude for InitializeRequest")
	envVarsPtr := flag.String("vars", "KEY=VALUE", "Enviornment variables in format: 'Key=Value, Key=Value'")
	profilePtr := flag.String("profile", "", "Profile JSON string")
	flag.Parse()

	// Set Enviornment variables
	if envVarsPtr != nil {
		golog.Info("Setting Enviornment variables.")

		// Split String Values
		keyValuePairs := strings.Split(*envVarsPtr, ",")

		// Iterate over keyValuePairs
		for _, v := range keyValuePairs {
			// Trim White Space
			tv := strings.TrimSpace(v)

			// Split Key Value Pairs
			value := strings.Split(tv, "=")
			if len(value) != 2 {
				return nil, errors.New("Invalid Enviornment Variable Format")
			}

			// Set Env Variables
			os.Setenv(value[0], value[1])
		}
	}

	// Set Location
	req := &api.InitializeRequest{
		Location: &common.Location{
			Latitude:  *latPtr,
			Longitude: *lngPtr,
		},
		Profile: common.NewDefaultProfile(),
	}

	// Set Profile
	if *profilePtr != "" {
		golog.Info("Setting Profile from JSON.")
		p := &common.Profile{}

		// Unmarshal JSON String
		err := protojson.Unmarshal([]byte(*profilePtr), p)
		if err == nil {
			req.Profile = p
		} else {
			golog.Warn("Failed to set Profile from flag")
		}
	}
	return req, nil
}
