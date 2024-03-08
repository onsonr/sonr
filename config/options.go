package config

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func persistentBanner(address string) string {
	return fmt.Sprintf(`
Sonr Highway
· Gateway: http://%s
· Node RPC: http://localhost:26657
`, address)
}

// HighwayOption is a function that sets some option on the HighwayOptions
type HighwayOption func(*Highway)

func (o *Highway) ListenAddress() string {
	return fmt.Sprintf("%s:%d", o.Host, o.GatewayPort)
}

// PrintBanner prints the banner
func (o *Highway) PrintBanner() {
	pterm.DefaultHeader.Printf(persistentBanner(fmt.Sprintf("localhost:%d", o.GatewayPort)))
}

// Serve starts the highway server
func (o *Highway) Serve(e *echo.Echo) {
	o.PrintBanner()
	e.Logger.Fatal(e.Start(o.ListenAddress()))
}

// NewHway returns a new HighwayOptions
<<<<<<<< HEAD:cmd/hway/config/options.go
func NewHway() *Highway {
	v := viper.New()
	v.SetEnvPrefix("HWAY")
	v.AutomaticEnv()
	conf := &Highway{
========
func NewHway() *HighwayConfig {
	return &HighwayConfig{
>>>>>>>> master:config/options.go
		GatewayPort: 8000,
		Host:        "0.0.0.0",
	}
	if err := v.Unmarshal(conf); err != nil {
		panic(err)
	}
	return conf
}

func (o *Highway) ReadFlags(c *cobra.Command) error {
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

	psql, err := c.Flags().GetString("hway-psql")
	if err != nil {
		return err
	}
	o.PostgresConnection = psql

	redis, err := c.Flags().GetString("hway-redis")
	if err != nil {
		return err
	}
	o.RedisConnection = redis
<<<<<<<< HEAD:cmd/hway/config/options.go

========
>>>>>>>> master:config/options.go
	return nil
}

// HasPostgres returns true if the postgres connection is set
<<<<<<<< HEAD:cmd/hway/config/options.go
func (o *Highway) HasPostgres() bool {
========
func (o *HighwayConfig) HasPostgres() bool {
>>>>>>>> master:config/options.go
	return o.PostgresConnection != ""
}

// HasRedis returns true if the redis connection is set
<<<<<<<< HEAD:cmd/hway/config/options.go
func (o *Highway) HasRedis() bool {
========
func (o *HighwayConfig) HasRedis() bool {
>>>>>>>> master:config/options.go
	return o.RedisConnection != ""
}

// Validate validates the HighwayOptions
func (o *Highway) Validate() error {
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
	return func(o *Highway) {
		o.GatewayPort = port
	}
}

// WithHost sets the Host
func WithHost(host string) HighwayOption {
	return func(o *Highway) {
		o.Host = host
	}
}
