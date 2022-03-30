package config

import (
	"errors"
	"os"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/kataras/golog"
)

var (
	// deviceID is the device ID. Either provided or found
	deviceID string

	// hostName is the host name. Either provided or found
	hostName string
)

var (
	logger = golog.Default.Child("internal/device")

	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("Duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("Prefix or Suffix set with Replace.")
	ErrSeparatorLength            = errors.New("Separator length must be 1.")
	ErrNoFileNameSet              = errors.New("File name was not set by options.")

	// Device ID Errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")

	// Node Errors
	ErrEmptyQueue       = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery     = errors.New("No SName or PeerID provided.")
	ErrMissingParam     = errors.New("Paramater is missing.")
	ErrProtocolsNotSet  = errors.New("Node Protocol has not been initialized.")
	ErrRoutingNotSet    = errors.New("DHT and Host have not been set by Routing Function")
	ErrListenerRequired = errors.New("Listener was not Provided")
	ErrMDNSInvalidConn  = errors.New("Invalid Connection, cannot begin MDNS Service")

	// Default P2P Properties
	BootstrapAddrStrs = []string{
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}
	AddrStoreTTL = time.Minute * 5
)

type Option func(o *options)

// SetDeviceID sets the device ID
func SetDeviceID(id string) Option {
	return func(o *options) {
		// Set Home Directory
		if id != "" {
			o.deviceID = id
		}
	}
}

// WithHomePath sets the Home Directory
func WithHomePath(p string) Option {
	return func(o *options) {
		// Set Home Directory
		if p != "" {
			o.HomeDir = p
		}
	}
}

// WithTempPath sets the Temporary Directory
func WithTempPath(p string) Option {
	return func(o *options) {
		// Set Home Directory
		if p != "" {
			o.TempDir = p
		}
	}
}

// WithSupportPath sets the Support Directory
func WithSupportPath(p string) Option {
	return func(o *options) {
		// Set Home Directory
		if p != "" {
			o.SupportDir = p
		}
	}
}

// options holds directory list
type options struct {
	HomeDir    string
	TempDir    string
	SupportDir string

	walletDir    string
	databaseDir  string
	downloadsDir string
	textileDir   string
	deviceID     string
}

// defaultOptions returns fsOptions
func defaultOptions() *options {
	opts := &options{}
	if IsDesktop() {
		hp, err := os.UserHomeDir()
		if err != nil {
			logger.Errorf("%s - Failed to get HomeDir, ", err)
		} else {
			opts.HomeDir = hp
		}

		tp, err := os.UserCacheDir()
		if err != nil {
			logger.Errorf("%s - Failed to get TempDir, ", err)
		} else {
			opts.TempDir = tp
		}

		sp, err := os.UserConfigDir()
		if err != nil {
			logger.Errorf("%s - Failed to get SupportDir, ", err)
		} else {
			opts.SupportDir = sp
		}

		id, err := machineid.ID()
		if err != nil {
			logger.Errorf("%s - Failed to get Device ID", err)
		} else {
			opts.deviceID = id
		}
	}
	return opts
}
