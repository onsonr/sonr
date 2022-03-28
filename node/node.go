package node

import (
	"context"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p"
	cmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/config"
	"github.com/spf13/viper"
	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
	"google.golang.org/protobuf/proto"
)

var (
	logger = golog.Default.Child("core/node")
	ctx    context.Context
)

// NewMotor Creates a node with its implemented protocols
func NewMotor(ctx context.Context, l net.Listener, options ...Option) (HostImpl, error) {
	// Initialize DHT
	opts := defaultOptions(config.Role_MOTOR)
	node, err := opts.Apply(ctx, options...)
	if err != nil {
		return nil, err
	}

	// Start Host
	node.Host, err = libp2p.New(ctx,
		libp2p.Identity(node.privKey),
		libp2p.ConnectionManager(cmgr.NewConnManager(
			opts.LowWater,    // Lowwater
			opts.HighWater,   // HighWater,
			opts.GracePeriod, // GracePeriod
		)),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(node.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		logger.Errorf("%s - NewHost: Failed to create libp2p host", err)
		return nil, err
	}
	node.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := node.Bootstrap(context.Background()); err != nil {
		logger.Errorf("%s - Failed to Bootstrap KDHT to Host", err)
		node.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range opts.BootstrapPeers {
		if err := node.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := node.createDHTDiscovery(opts); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		node.SetStatus(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for MDNS
	node.createMdnsDiscovery(opts)
	node.SetStatus(Status_READY)
	go node.Serve()

	// Open Store with profileBuf
	return node, nil
}

// NewHighway Creates a node with its implemented protocols
func NewHighway(ctx context.Context, options ...Option) {

	// Initialize DHT
	opts := defaultOptions(config.Role_HIGHWAY)
	node, err := opts.Apply(ctx, options...)
	if err != nil {
		panic(err)
	}

	// Open Listener on Port
	l, err := net.Listen(opts.network, opts.Address())
	if err != nil {
		golog.Default.Child("(app)").Fatalf("%s - Failed to Create New Listener", err)
		panic(err)
	}

	// Start Host
	node.Host, err = libp2p.New(ctx,
		libp2p.Identity(node.privKey),
		libp2p.ConnectionManager(cmgr.NewConnManager(
			opts.LowWater,    // Lowwater
			opts.HighWater,   // HighWater,
			opts.GracePeriod, // GracePeriod
		)),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(node.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		logger.Errorf("%s - NewHost: Failed to create libp2p host", err)
		panic(err)
	}
	node.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := node.Bootstrap(context.Background()); err != nil {
		logger.Errorf("%s - Failed to Bootstrap KDHT to Host", err)
		node.SetStatus(Status_FAIL)
		panic(err)
	}

	// Connect to Bootstrap Nodes
	for _, pi := range opts.BootstrapPeers {
		if err := node.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := node.createDHTDiscovery(opts); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		node.SetStatus(Status_FAIL)
		panic(err)
	}

	// Initialize Discovery for MDNS
	node.createMdnsDiscovery(opts)
	node.SetStatus(Status_READY)
	go node.Serve()
	persist(l)
}

// HostID returns the ID of the Host
func (n *node) HostID() peer.ID {
	return n.Host.ID()
}

// Ping sends a ping to the peer
func (n *node) Ping(pid string) error {
	return nil
}

// Publish publishes a message to the network
func (n *node) Publish(t string, message proto.Message, metadata *types.Metadata) error {
	return nil
}

// Role returns the role of the node
func (n *node) Role() config.Role {
	return n.mode
}

// persist contains the main loop for the Node
func persist(l net.Listener) {
	golog.Default.Child("(app)").Infof("Starting GRPC Server on %s", l.Addr().String())
	// Check if CLI Mode
	if config.IsMobile() {
		golog.Default.Child("(app)").Info("Skipping Serve, Node is mobile...")
		return
	}

	// Wait for Exit Signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exit(0)
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

// Exit handles cleanup on Sonr Node
func exit(code int) {
	golog.Default.Child("(app)").Debug("Cleaning up Node on Exit...")
	defer ctx.Done()

	// Check for Full Desktop Node
	if config.IsDesktop() {
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
