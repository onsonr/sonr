package config

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func persistentBanner(address string) string {
	return fmt.Sprintf(`
Sonr Highway
· Gateway: http://%s
· Node RPC: http://localhost:26657
`, address)
}

// HighwayOption is a function that sets some option on the HighwayOptions
type HighwayOption func(*HighwayConfig)

func (o *HighwayConfig) ListenAddress() string {
	return fmt.Sprintf("%s:%d", o.Host, o.GatewayPort)
}

// PrintBanner prints the banner
func (o *HighwayConfig) PrintBanner() {
	if o.EnableBanner {
		pterm.DefaultHeader.Printf(persistentBanner(fmt.Sprintf("localhost:%d", o.GatewayPort)))
	}
}

// Serve starts the highway server
func (o *HighwayConfig) Serve(e *echo.Echo) {
	o.PrintBanner()
	e.Logger.Fatal(e.Start(o.ListenAddress()))
}

// CreateHwayConfig returns a new HighwayOptions
func CreateHwayConfig() *HighwayConfig {
	return &HighwayConfig{
		GatewayPort:  8000,
		Host:         "0.0.0.0",
		EnableBanner: true,
	}
}

func (o *HighwayConfig) ReadFlags(c *cobra.Command) error {
	host, err := c.Flags().GetString("hway-host")
	if err != nil {
		return err
	}
	o.Host = host

	port, err := c.Flags().GetInt("hway-port")
	if err != nil {
		return err
	}
	o.GatewayPort = port
	return nil
}

// Validate validates the HighwayOptions
func (o *HighwayConfig) Validate() error {
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
	return func(o *HighwayConfig) {
		o.GatewayPort = port
	}
}

// WithHost sets the Host
func WithHost(host string) HighwayOption {
	return func(o *HighwayConfig) {
		o.Host = host
	}
}

// WithEnableBanner sets the EnableBanner
func WithEnableBanner(enable bool) HighwayOption {
	return func(o *HighwayConfig) {
		o.EnableBanner = enable
	}
}
