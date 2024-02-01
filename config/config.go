package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Config is the configuration for the application
type Config struct {
	Name            string `mapstructure:"name" json:"name"`
	Version         string `mapstructure:"version" json:"version"`
	Description     string `mapstructure:"description" json:"description"`
	DefaultNodeHome string `mapstructure:"default_node_home" json:"default_node_home"`
	AppConfigYAML   []byte `mapstructure:"app_config_yaml" json:"app_config_yaml"`
	HighwayEnabled  bool   `mapstructure:"highway_enabled" json:"highway_enabled"`
}

// LoadConfig loads the configuration from the file
func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "mapstructure"
	}); err != nil {
		return nil, err
	}

	return &config, nil
}
