package config

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
)

func persistentBanner(address string) string {
	return fmt.Sprintf(`
Sonr Highway
· Gateway: http://%s
· Node RPC: http://localhost:26657
`, address)
}

// HighwayOption is a function that sets some option on the HighwayOptions
type HighwayOption func(*HighwayOptions)

// HighwayOptions holds the options for the highway server
type HighwayOptions struct {
	// GatewayPort is the port that the gateway listens on
	GatewayPort int `json:"gateway_port"`

	// Host is the host that the gateway listens on
	Host string `json:"host"`

	// EnableBanner enables the banner
	EnableBanner bool `json:"enable_banner"`
}

func (o *HighwayOptions) ListenAddress() string {
	return fmt.Sprintf("%s:%d", o.Host, o.GatewayPort)
}

// PrintBanner prints the banner
func (o *HighwayOptions) PrintBanner() {
	if o.EnableBanner {
		pterm.DefaultHeader.Printf(persistentBanner(fmt.Sprintf("localhost:%d", o.GatewayPort)))
	}
}

// Serve starts the highway server
func (o *HighwayOptions) Serve(e *echo.Echo) {
	o.PrintBanner()
	e.Logger.Fatal(e.Start(o.ListenAddress()))
}

// NewHighwayOptions returns a new HighwayOptions
func NewHighwayOptions() *HighwayOptions {
	return &HighwayOptions{
		GatewayPort:  8000,
		Host:         "0.0.0.0",
		EnableBanner: true,
	}
}

// Validate validates the HighwayOptions
func (o *HighwayOptions) Validate() error {
	if o.GatewayPort < 0 {
		return fmt.Errorf("gateway port must be greater than 0")
	}
	if o.Host == "" {
		return fmt.Errorf("host must not be empty")
	}
	return nil
}

// WithGatewayPort sets the GatewayPort
func WithGatewayPort(port int) HighwayOption {
	return func(o *HighwayOptions) {
		o.GatewayPort = port
	}
}

// WithHost sets the Host
func WithHost(host string) HighwayOption {
	return func(o *HighwayOptions) {
		o.Host = host
	}
}

// WithEnableBanner sets the EnableBanner
func WithEnableBanner(enable bool) HighwayOption {
	return func(o *HighwayOptions) {
		o.EnableBanner = enable
	}
}
