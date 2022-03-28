package config

import (
	"errors"
	"os"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

var (
	// Determined/Provided Paths
	Home      Folder // ApplicationDocumentsDir on Mobile, HOME_DIR on Desktop
	Support   Folder // AppSupport Directory
	Temporary Folder // AppCache Directory

	// Calculated Paths
	Database   Folder // Device DB Folder
	Downloads  Folder // Temporary Directory on Mobile for Export, Downloads on Desktop
	Wallet     Folder // Encrypted Storage Directory
	ThirdParty Folder // Sub-Directory of Support, used for Textile

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

// Apply sets device directories for Path
func (fo *options) Apply() error {
	// Get the hostname
	hn, err := os.Hostname()
	if err != nil {
		logger.Errorf("%s - Failed to get HostName", err)
		return err
	}
	hostName = hn

	// Check if deviceID is set
	if fo.deviceID == "" {
		logger.Errorf("%s - Device ID is empty", ErrEmptyDeviceID)
		return ErrEmptyDeviceID
	}
	deviceID = fo.deviceID

	// Check for Valid
	if fo.HomeDir == "" {
		return errors.New("Home Directory was not set.")
	}
	if fo.SupportDir == "" {
		return errors.New("Support Directory was not set.")
	}
	if fo.TempDir == "" {
		return errors.New("Temporary Directory was not set.")
	}

	// Set Home Folder
	Home = Folder(fo.HomeDir)
	Support = Folder(fo.SupportDir)
	Temporary = Folder(fo.TempDir)

	// Create Downloads Folder
	if IsDesktop() {
		Downloads, err = Home.CreateFolder("Downloads")
		if err != nil {
			return err
		}
	} else {
		Downloads, err = Temporary.CreateFolder("Downloads")
		if err != nil {
			return err
		}
	}

	// Create Database Folder
	Database, err = Support.CreateFolder(".db")
	if err != nil {
		return err
	}

	// Create Wallet Folder
	Wallet, err = Support.CreateFolder(".wallet")
	if err != nil {
		return err
	}

	// Create Third Party Folder
	ThirdParty, err = Support.CreateFolder("third_party")
	if err != nil {
		return err
	}
	return nil
}

type SonrConfig struct {
	HighwayAddress       string   `json:"highway_address"`
	HighwayPort          int      `json:"highway_port"`
	HighwayNetwork       string   `json:"highway_network"`
	HighwayDID           string   `json:"highway_did"`
	IPFSPort             int      `json:"ipfs_port"`
	IPFSPath             string   `json:"ipfs_path"`
	LibP2PLowWater       int      `json:"libp2p_low_water"`
	LibP2PHighWater      int      `json:"libp2p_high_water"`
	LibP2PRendevouz      string   `json:"libp2p_rendevouz"`
	LibP2PBootstrapPeers []string `json:"libp2p_bootstrap_peers"`
	HomeDir              string   `json:"home_dir"`
	CacheDir             string   `json:"cache_dir"`
	ConfigDir            string   `json:"config_dir"`
	WalletDir            string   `json:"wallet_dir"`
	DeviceId             string   `json:"device_id"`
	PublicIP             string   `json:"public_ip"`
	PrivateIP            string   `json:"private_ip"`
	AccountName          string   `json:"account_name"`
}

func (sc *SonrConfig) Save() (*SonrConfig, error) {
	viper.Set("highway.address", sc.HighwayAddress)
	viper.Set("highway.port", sc.HighwayPort)
	viper.Set("highway.network", sc.HighwayNetwork)
	viper.Set("highway.did", sc.HighwayDID)
	viper.Set("ipfs.port", sc.IPFSPort)
	viper.Set("ipfs.path", sc.IPFSPath)
	viper.Set("libp2p.lowWater", sc.LibP2PLowWater)
	viper.Set("libp2p.highWater", sc.LibP2PHighWater)
	viper.Set("libp2p.rendevouz", sc.LibP2PRendevouz)
	viper.Set("libp2p.bootstrap_peers", sc.LibP2PBootstrapPeers)
	viper.Set("home_dir", sc.HomeDir)
	viper.Set("cache_dir", sc.CacheDir)
	viper.Set("config_dir", sc.ConfigDir)
	viper.Set("wallet_dir", sc.WalletDir)
	viper.Set("device_id", sc.DeviceId)
	viper.Set("public_ip", sc.PublicIP)
	viper.Set("private_ip", sc.PrivateIP)
	viper.Set("account_name", sc.AccountName)
	err := viper.WriteConfig()
	if err != nil {
		return nil, err
	}
	return sc, nil
}

// Return the config dir path as a Folder.
func (sc *SonrConfig) ConfigFolder() Folder {
	return Folder(sc.ConfigDir)
}

// Return the home dir path as a Folder.
func (sc *SonrConfig) HomeFolder() Folder {
	return Folder(sc.HomeDir)
}

// Return the cache dir path as a Folder.
func (sc *SonrConfig) CacheFolder() Folder {
	return Folder(sc.CacheDir)
}

// Create or return the wallet directory as a Folder.
func (sc *SonrConfig) WalletFolder() Folder {
	return Folder(sc.WalletDir)
}
