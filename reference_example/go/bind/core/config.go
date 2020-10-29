package core

import (
	"encoding/json"
	"io/ioutil"

	ipfs_config "github.com/ipfs/go-ipfs-config"
	ipfs_common "github.com/ipfs/go-ipfs/repo/common"
)

type Config struct {
	cfg *ipfs_config.Config
}

func NewConfig(raw_json []byte) (cfg *Config, err error) {
	cfg = &Config{}
	err = cfg.Set(raw_json)
	return cfg, err
}

func NewDefaultConfig() (*Config, error) {
	cfg, err := initConfig(ioutil.Discard, 2048)
	if err != nil {
		return nil, err
	}

	return &Config{cfg}, nil
}

func (c *Config) getConfig() (cfg *ipfs_config.Config) {
	return c.cfg
}

func (c *Config) Set(raw_json []byte) (err error) {
	var mapcfg map[string]interface{}
	if err = json.Unmarshal(raw_json, &mapcfg); err != nil {
		return
	}

	var cfg *ipfs_config.Config

	if cfg, err = ipfs_config.FromMap(mapcfg); err == nil {
		c.cfg = cfg
	}

	return
}

func (c *Config) Get() ([]byte, error) {
	return json.Marshal(c.cfg)
}

func (c *Config) SetKey(key string, raw_value []byte) error {
	var val interface{}

	if err := json.Unmarshal(raw_value, &val); err != nil {
		return err
	}

	cfg, err := ipfs_config.ToMap(c.cfg)
	if err != nil {
		return err
	}

	if err := ipfs_common.MapSetKV(cfg, key, val); err != nil {
		return err
	}

	// update Config
	newcfg, err := ipfs_config.FromMap(cfg)
	if err == nil {
		c.cfg = newcfg
	}
	return err
}

func (c *Config) GetKey(key string) ([]byte, error) {
	cfg, err := ipfs_config.ToMap(c.cfg)
	if err != nil {
		return nil, err
	}

	val, err := ipfs_common.MapGetKV(cfg, key)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&val)
}
