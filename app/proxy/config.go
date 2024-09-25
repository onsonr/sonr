package proxy

import "github.com/spf13/viper"

type Config struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	SSL      bool   `mapstructure:"SSL"`
	CertFile string `mapstructure:"CERT_FILE"`
	KeyFile  string `mapstructure:"KEY_FILE"`
}

func (c *Config) GetHostname() string {
	return c.Host + ":" + c.Port
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
