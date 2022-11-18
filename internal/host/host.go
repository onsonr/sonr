package host

// mplex "github.com/libp2p/go-libp2p-mplex"

/// direct "github.com/libp2p/go-libp2p-webrtc-direct"

// "github.com/pion/webrtc/v3"

// // hostImpl type - a p2p host implementing one or more p2p protocols
// type hostImpl struct {
// 	// Standard Node Implementation
// 	host    host.Host
// 	accAddr string

// 	// Host and context
// 	privKey     crypto.PrivKey
// 	dhtPeerChan <-chan peer.AddrInfo

// 	// Properties
// 	ctx context.Context

// 	*dht.IpfsDHT
// 	*ps.PubSub

// 	// State
// 	fsm *SFSM
// }

// // NewDefaultHost Creates a Sonr libp2p Host with the given config
// func NewDefaultHost(ctx context.Context) (SonrHost, error) {
// 	var err error
// 	// Create the host.
// 	hn := &hostImpl{
// 		ctx: ctx,
// 		fsm: NewFSM(ctx),
// 	}
// 	// findPrivKey returns the private key for the host.
// 	findPrivKey := func() (crypto.PrivKey, error) {
// 		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
// 		if err == nil {

// 			return privKey, nil
// 		}
// 		return nil, err
// 	}
// 	// Fetch the private key.
// 	hn.privKey, err = findPrivKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create Connection Manager
// 	cnnmgr, err := cmgr.NewConnManager(c.Libp2pLowWater, c.Libp2pHighWater)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Start Host
// 	hn.host, err = libp2p.New(
// 		libp2p.Identity(hn.privKey),
// 		libp2p.ConnectionManager(cnnmgr),
// 		libp2p.DefaultListenAddrs,
// 		libp2p.Routing(hn.Router),
// 		libp2p.EnableAutoRelay(),
// 	)
// 	if err != nil {
// 		hn.fsm.SetState(Status_FAIL)
// 		return nil, err
// 	}
// 	hn.fsm.SetState(Status_CONNECTING)

// 	// Bootstrap DHT
// 	if err := hn.Bootstrap(context.Background()); err != nil {
// 		hn.fsm.SetState(Status_FAIL)
// 		return nil, err
// 	}

// 	// Connect to Bootstrap Nodes
// 	for _, pi := range c.Libp2pBootstrapPeers {
// 		if err := hn.Connect(pi); err != nil {
// 			continue
// 		} else {
// 			hn.fsm.SetState(Status_FAIL)

// 			break
// 		}
// 	}

// 	// Initialize Discovery for DHT
// 	if err := hn.createDHTDiscovery(c); err != nil {
// 		// Check if we need to close the listener
// 		hn.fsm.SetState(Status_FAIL)
// 		return nil, err
// 	}

// 	hn.fsm.SetState(Status_READY)
// 	go hn.Serve()
// 	return hn, nil
// }

// // Address returns the address of the underlying wallet
// func (h *hostImpl) Address() string {
// 	return h.accAddr
// }

// // createDHTDiscovery is a Helper Method to initialize the DHT Discovery
// func (hn *hostImpl) createDHTDiscovery() error {
// 	// Set Routing Discovery, Find Peers
// 	var err error
// 	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
// 	// dsc.Advertise(hn.ctx, routingDiscovery, c.Libp2pRendezvous, c.Libp2pTTL)

// 	// Create Pub Sub
// 	hn.PubSub, err = ps.NewGossipSub(hn.ctx, hn.host, ps.WithDiscovery(routingDiscovery))
// 	if err != nil {
// 		hn.fsm.SetState(Status_FAIL)
// 		return err
// 	}

// 	// Handle DHT Peers
// 	// hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, c.Libp2pRendezvous, c.Libp2pTTL)
// 	if err != nil {
// 		hn.fsm.SetState(Status_FAIL)
// 		return err
// 	}

// 	hn.fsm.SetState(Status_READY)
// 	return nil
// }

// func (hn *hostImpl) Close() error {
// 	err := hn.host.Close()
// 	if err != nil {
// 		hn.fsm.SetState(Status_FAIL)
// 		return err
// 	}

// 	hn.fsm.SetState(Status_STANDBY)

// 	return nil
// }

// // NeedsWait checks if state is Resumed or Paused and blocks channel if needed
// func (hn *hostImpl) NeedsWait() {
// 	<-hn.fsm.Chn
// }

// /*
// Stops the libp2p host, dhcp, and sets the host status to IDLE
// */
// func (hn *hostImpl) Stop() error {
// 	err := hn.host.Close()
// 	if err != nil {
// 		hn.fsm.SetState(Status_FAIL)
// 		return err
// 	}
// 	hn.Pause()

// 	return nil
// }

// /*
// Stops the libp2p host, dhcp, and sets the host status to ready
// */
// func (hn *hostImpl) Pause() error {
// 	defer hn.fsm.PauseOperation()
// 	hn.fsm.SetState(Status_STANDBY)
// 	return nil
// }

// func (hn *hostImpl) Resume() error {
// 	defer hn.fsm.ResumeOperation()
// 	hn.fsm.SetState(Status_STANDBY)

// 	return nil
// }

// func (hn *hostImpl) Status() HostStatus {
// 	return hn.fsm.CurrentStatus
// }

// // Peer is a Helper Method to get the peer from the host
// func (hn *hostImpl) Peer() (*ct.Peer, error) {
// 	return &ct.Peer{
// 		PeerId: hn.host.ID().String(),
// 		Did:    addrToDidUrl(hn.accAddr),
// 	}, nil
// }
