package config

import (
	"fmt"
	"os"
	"runtime"
	"time"

	mv1 "go.buf.build/grpc/go/sonr-io/core/motor/v1"
	t "go.buf.build/grpc/go/sonr-io/core/types/v1"
)

// CallbackImpl is the implementation of Callback interface
type CallbackImpl interface {
	// OnRefresh is called when the LobbyProtocol is refreshed and pushes a RefreshEvent
	OnRefresh(event *mv1.OnLobbyRefreshResponse)

	// OnMailbox is called when the MailboxProtocol receives a MailboxEvent
	OnMailbox(event *mv1.OnMailboxMessageResponse)

	// OnInvite is called when the TransferProtocol receives InviteEvent
	OnInvite(event *mv1.OnTransmitInviteResponse)

	// OnDecision is called when the TransferProtocol receives a DecisionEvent
	OnDecision(event *mv1.OnTransmitDecisionResponse, invite *mv1.OnTransmitInviteResponse)

	// OnProgress is called when the TransferProtocol sends or receives a ProgressEvent
	OnProgress(event *mv1.OnTransmitProgressResponse)

	// OnTransfer is called when the TransferProtocol completes a transfer and pushes a CompleteEvent
	OnComplete(event *mv1.OnTransmitCompleteResponse)
}

// Role is the type of the node (Client, Highway)
type Role int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	Role_UNSPECIFIED Role = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	Role_TEST

	// Role_MOTOR is for a Motor Node
	Role_MOTOR

	// Role_HIGHWAY is for a Highway Node
	Role_HIGHWAY
)

// Motor returns true if the node has a client.
func (m Role) IsMotor() bool {
	return m == Role_MOTOR
}

// Highway returns true if the node has a highway stub.
func (m Role) IsHighway() bool {
	return m == Role_HIGHWAY
}

// Prefix returns golog prefix for the node.
func (m Role) Prefix() string {
	var name string
	switch m {
	case Role_HIGHWAY:
		name = "highway"
	case Role_MOTOR:
		name = "motor"
	case Role_TEST:
		name = "test"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}

type Configuration struct {
	connection       t.Connection
	deviceId         string
	location         *t.Location
	profile          *t.Profile
	homeDirectory    string
	supportDirectory string
	tempDirectory    string
}

func DefaultConfiguration() *Configuration {
	// Default configuration
	c := &Configuration{
		connection: t.Connection_CONNECTION_WIFI,
		location:   NewDefaultLocation(),
		profile:    NewDefaultProfile(),
	}

	// Check for non-mobile device
	if !IsMobile() {

		// Set Device ID
		if fid, err := ID(); err == nil {
			c.deviceId = fid
		}

		// Set Home Directory
		if hdir, err := os.UserHomeDir(); err == nil {
			c.homeDirectory = hdir
		}

		// Set Support Directory
		if sdir, err := os.UserConfigDir(); err == nil {
			c.supportDirectory = sdir
		}

		// Set Temp Directory
		if tdir, err := os.UserCacheDir(); err == nil {
			c.tempDirectory = tdir
		}
	}
	return c
}

// NewDefaultLocation returns the Sonr HQ as default location
func NewDefaultLocation() *t.Location {
	return &t.Location{
		Latitude:  float64(40.673010),
		Longitude: float64(-73.994450),
		Placemark: &t.Location_Placemark{
			Name:                  "Sonr HQ",
			Street:                "94 9th St.",
			IsoCountryCode:        "US",
			Country:               "United States",
			AdministrativeArea:    "New York",
			SubAdministrativeArea: "Brooklyn",
			Locality:              "Brooklyn",
			SubLocality:           "Gowanus",
			PostalCode:            "11215",
		},
	}
}

// DefaultProfileOption is a type for Profile Options
type DefaultProfileOption func(profileOpts)

// profileOpts contains the options for generating a new Profile
type profileOpts struct {
	refProfile *t.Profile
	sname      string
	firstName  string
	lastName   string
	picture    []byte
	bio        string
}

// defaultProfileOpts returns the default profile options
func defaultProfileOpts() profileOpts {
	return profileOpts{
		sname:     fmt.Sprintf("a%s", runtime.GOOS),
		firstName: "Anonymous",
		lastName:  runtime.GOOS,
		picture:   make([]byte, 0),
	}
}

// NewDefaultProfile creates a new default Profile
func NewDefaultProfile(options ...DefaultProfileOption) *t.Profile {
	// Set default options
	opts := defaultProfileOpts()
	for _, option := range options {
		option(opts)
	}

	// Create Profile Build Func
	buildProfile := func(opts profileOpts) *t.Profile {
		return &t.Profile{
			SName:        opts.sname,
			FirstName:    opts.firstName,
			LastName:     opts.lastName,
			Picture:      opts.picture,
			Bio:          opts.bio,
			LastModified: time.Now().Unix(),
		}
	}

	// Check if refProfile is set and build
	if opts.refProfile != nil {
		if !checkProfile(opts.refProfile) {
			return buildProfile(opts)
		}
		return opts.refProfile
	}
	return buildProfile(opts)
}

// WithCheckerProfile sets the checker profile
func WithCheckerProfile(profile *t.Profile) DefaultProfileOption {
	return func(opts profileOpts) {
		opts.refProfile = profile
	}
}

// checkProfile checks if the Profile is valid
func checkProfile(p *t.Profile) bool {
	if p == nil {
		return false
	}
	if len(p.SName) == 0 || len(p.FirstName) == 0 {
		return false
	}
	return true
}
