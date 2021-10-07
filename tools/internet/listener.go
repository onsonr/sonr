package internet

import (
	"context"
	"fmt"
	"strings"
	"time"

	"net"

	"github.com/kataras/golog"
	reuse "github.com/libp2p/go-reuseport"
	ma "github.com/multiformats/go-multiaddr"
)

const (
	DefaultListenNetwork = "tcp"
	DefaultListenIP      = "127.0.0.1"
	DefaultListenHost    = "localhost"
	DefaultListenPort    = 0
)

// TCPListenerOpt is a function that modifies the tcp listener options.
type TCPListenerOpt func(tcpOptions)

// WithPort sets the port of the listener
func WithPort(p int) TCPListenerOpt {
	// defaultPortFunc returns TCPListenerOpt that sets port to default
	defaultPortFunc := func() TCPListenerOpt {
		return func(o tcpOptions) {
			o.port = DefaultListenPort
		}
	}

	// setPortFunc returns TCPListenerOpt that sets port to provided
	setPortFunc := func(port int) TCPListenerOpt {
		return func(o tcpOptions) {
			o.port = port
		}
	}

	// Check if port is valid and return defaultPortFunc if not
	results := checkPorts(DefaultListenHost, p)
	if results[fmt.Sprint(p)] == "failed" {
		logger.Child("internet/TCPListener").Error(fmt.Errorf("port %v is not available", p))
		return defaultPortFunc()
	}
	return setPortFunc(p)
}

// WithHost sets the bootstrap peers.
func WithHost(h string) TCPListenerOpt {
	return func(o tcpOptions) {
		if h != "" {
			o.host = h
		}
	}
}

// tcpOptions is the options for the TCP listener
type tcpOptions struct {
	network string
	host    string
	port    int
}

// defaultTCPListenerOpts returns the default options for the TCP listener
func defaultTCPListenerOpts() tcpOptions {
	opts := tcpOptions{
		network: DefaultListenNetwork,
		host:    DefaultListenHost,
		port:    DefaultListenPort,
	}

	return opts
}

// Apply applies the options to the listener
func (o tcpOptions) Apply(options ...TCPListenerOpt) {
	if len(options) > 0 {
		logger.Child("internet/TCPListener").Info(fmt.Sprintf("Applying %v options to TCPListener: %s", len(options), o.getAddress()))
		for _, opt := range options {
			opt(o)
		}
	} else {
		logger.Child("internet/TCPListener").Info(fmt.Sprintf("Default Options set for TCPListener: %s", o.getAddress()))
	}
}

// getAddress returns the address of the listener
func (o tcpOptions) getAddress() string {
	return fmt.Sprintf("%s:%v", o.host, o.port)
}

// TCPListener is a cross-platform Listener
type TCPListener struct {
	ctx context.Context
	net.Listener
	opts tcpOptions
}

// NewTCPListener Creates a new TCP Listener
func NewTCPListener(ctx context.Context, options ...TCPListenerOpt) (*TCPListener, error) {
	opts := defaultTCPListenerOpts()
	opts.Apply(options...)

	// Open Listener
	lc, err := reuse.Listen(opts.network, opts.getAddress())
	if err != nil {
		return nil, err
	}

	// Create Listener
	go handlePort(ctx, lc)
	return &TCPListener{
		ctx:      ctx,
		Listener: lc,
		opts:     opts,
	}, nil
}

// IsIPv4 returns true if the listener is an IPv4 listener
func (l *TCPListener) IsIPv4() bool {
	return l.Addr().(*net.TCPAddr).IP.To4() != nil
}

// IsIPv6 returns true if the listener is an IPv6 listener
func (l *TCPListener) IsIPv6() bool {
	return l.Addr().(*net.TCPAddr).IP.To16() != nil
}

// Host returns the Host of the listener
func (l *TCPListener) Host() string {
	// Get Host
	hostStr := strings.Split(l.Addr().String(), ":")[0]
	hostStrSplits := strings.Split(hostStr, ".")

	// Validate IP address
	if len(hostStrSplits) != 4 {
		return DefaultListenIP
	}
	return hostStr
}

// Port returns the port of the listener
func (l *TCPListener) Port() int {
	return l.Addr().(*net.TCPAddr).Port
}

// Multiaddr returns the multiaddr of the listener
func (l *TCPListener) Multiaddr() (ma.Multiaddr, error) {
	return ma.NewMultiaddr(l.MultiaddrStr())
}

// MultiaddrStr returns the multiaddr string of the listener
func (l *TCPListener) MultiaddrStr() string {
	// Get Variables
	t := l.Transport()
	h := l.Host()
	n := l.Network()
	p := l.Port()

	// Logging and return
	maStr := fmt.Sprintf("/%s/%s/%s/%d", t, h, n, p)
	logger.Child("internet/TCPListener").Info("Created MultiAddr for TCPListener", golog.Fields{
		"Transport": t,
		"Host":      h,
		"Network":   n,
		"Port":      p,
		"MultiAddr": maStr,
	})
	return maStr
}

// Multiaddr returns the multiaddr of the listener
func (l *TCPListener) Network() string {
	if l.opts.network != "" {
		return l.opts.network
	}
	return DefaultListenNetwork
}

// Transport returns 'ip4' or 'ip6' based on the listener
func (l *TCPListener) Transport() string {
	if l.IsIPv6() {
		return "ip6"
	}
	return "ip4"
}

// handlePort handles the port of the listener until it is closed
func handlePort(ctx context.Context, l net.Listener) {
	for {
		select {
		case <-ctx.Done():
			err := l.Close()
			if err != nil {
				golog.Fatal("Failed to close Listening Port", err)
			}
			return
		}
	}
}

// checkPorts checks if the ports are available on the host
func checkPorts(ip string, ports ...int) map[string]string {
	// check emqx 1883, 8083 port
	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, fmt.Sprint(port))
		// 3 second timeout
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results[fmt.Sprint(port)] = "failed"
			continue
		} else {
			if conn != nil {
				results[fmt.Sprint(port)] = "success"
				_ = conn.Close()
			} else {
				results[fmt.Sprint(port)] = "failed"
			}
		}
	}
	return results
}
