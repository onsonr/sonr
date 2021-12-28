package node

import (
	"context"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/device"
	"github.com/sonr-io/core/types/go/node/motor/v1"
	"github.com/sonr-io/core/wallet"
	"github.com/spf13/viper"
)

// Start starts the Sonr Node
func Start(req *motor.InitializeRequest, options ...Option) {
	// Check if Node is already running
	if instance != nil {
		golog.Error("Sonr Instance already active")
		return
	}

	// Set Options
	opts := defaultOptions()
	for _, o := range options {
		o(opts)
	}

	// Set Logging Settings
	golog.SetLevel(opts.logLevel)
	golog.SetPrefix(opts.mode.Prefix())

	// Create Node
	ctx = context.Background()
	// Set Environment Variables
	vars := req.GetVariables()
	count := len(vars)

	// Iterate over Variables
	if count > 0 {
		for k, v := range vars {
			os.Setenv(k, v)
		}

		golog.Debug("Added Enviornment Variable(s)", golog.Fields{
			"Total": count,
		})
	}

	// Start File System
	device.Init(
		device.WithHomePath(req.GetDeviceOptions().GetHomeDir()),
		device.WithSupportPath(req.GetDeviceOptions().GetSupportDir()),
		device.SetDeviceID(req.GetDeviceOptions().GetId()),
	)

	// Open Keychain
	if wallet.Exists() {
		err := wallet.Open(wallet.WithPassphrase(req.GetWalletPassphrase()), wallet.WithSName(req.GetProfile().GetSName()))
		if err != nil {
			golog.Default.Child("(app)").Fatalf("%s - Failed to Open Keychain", err)
			Exit(1)
		}
	} else {
		err := wallet.New(req.GetWalletPassphrase(), req.GetProfile().GetSName())
		if err != nil {
			golog.Default.Child("(app)").Fatalf("%s - Failed to Create Keychain", err)
			Exit(1)
		}
	}
	// Open Listener on Port
	listener, err := net.Listen(opts.network, opts.Address())
	if err != nil {
		golog.Default.Child("(app)").Fatalf("%s - Failed to Create New Listener", err)
		return
	}

	// Set Node Stub
	instance, err = NewMotor(ctx, listener,
		WithMode(opts.mode))
	if err != nil {
		golog.Default.Child("(app)").Fatalf("%s - Failed to Start new Node", err)
	}

	// Serve listener for node grpc
	Persist(listener)
}

// Persist contains the main loop for the Node
func Persist(l net.Listener) {
	if instance == nil {
		golog.Error("Node instance is nil")
		return
	}

	golog.Default.Child("(app)").Infof("Starting GRPC Server on %s", l.Addr().String())
	// Check if CLI Mode
	if device.IsMobile() {
		golog.Default.Child("(app)").Info("Skipping Serve, Node is mobile...")
		return
	}

	// Wait for Exit Signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Exit(0)
	}()

	// Hold until Exit Signal
	for {
		select {
		case <-ctx.Done():
			golog.Default.Child("(app)").Info("Context Done")
			l.Close()
			return
		}
	}
}

// Pause calls the Pause function on the Node
func Pause() {
	if instance != nil {
		instance.Pause()
	}
}

// Resume calls the Resume function on the Node
func Resume() {
	if instance != nil {
		instance.Resume()
	}
}

// Exit handles cleanup on Sonr Node
func Exit(code int) {
	if instance == nil {
		golog.Default.Child("(app)").Debug("Skipping Exit, instance is nil...")
		return
	}
	golog.Default.Child("(app)").Debug("Cleaning up Node on Exit...")
	instance.Close()

	defer ctx.Done()

	// Check for Full Desktop Node
	if device.IsDesktop() {
		golog.Default.Child("(app)").Debug("Removing Bitcask DB...")
		ex, err := os.Executable()
		if err != nil {
			golog.Default.Child("(app)").Errorf("%s - Failed to find Executable", err)
			return
		}

		// Delete Executable Path
		exPath := filepath.Dir(ex)
		err = os.RemoveAll(filepath.Join(exPath, "sonr_bitcask"))
		if err != nil {
			golog.Default.Child("(app)").Warn("Failed to remove Bitcask, ", err)
		}
		err = viper.SafeWriteConfig()
		if err == nil {
			golog.Default.Child("(app)").Debug("Wrote new config file to Disk")
		}
		os.Exit(code)
	}
}
