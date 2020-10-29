package core

import (
	"encoding/json"
	"strings"
	"testing"

	ipfs_config "github.com/ipfs/go-ipfs-config"
)

const sampleFakeConfig = `
{
	"Addresses": {
		"API": "/ip4/127.0.0.1/tcp/5001",
		"Swarm": [
			"/ip4/0.0.0.0/tcp/0",
			"/ip6/::/tcp/0"
		]
	},
	"Bootstrap": [
		"/ip4/127.0.0.1/tcp/4001/ipfs/12D3KooWDWJ473M3fXMEcajbaGtqgr6i6SvDdh5Ru9i5ZzoJ9Qy8"
	]
}
`

func TestGetSetConfig(t *testing.T) {
	cfg, err := NewConfig([]byte(sampleFakeConfig))
	if err != nil {
		t.Fatal(err)
	}

	// get the whole config
	_, err = cfg.Get()
	if err != nil {
		t.Error(err)
	}

	// get a fake key
	val, err := cfg.GetKey("FAKEKEY")
	if err == nil {
		t.Errorf("exected no values but got: %s", val)
	}

	// get Api value
	val, err = cfg.GetKey("Addresses.API")
	if err != nil {
		t.Fatal(err)
	}

	var apiaddr string

	// check if api value is correct
	err = json.Unmarshal(val, &apiaddr)
	if err != nil {
		t.Error(err)
	}

	if apiaddr != "/ip4/127.0.0.1/tcp/5001" {
		t.Errorf("expected `/ip4/127.0.0.1/tcp/5001` got `%s`", apiaddr)
	}

	// get bootstrap value
	val, err = cfg.GetKey("Bootstrap")
	if err != nil {
		t.Fatal(err)
	}

	var bootstrapAddrs []string

	// check bootstrap value
	err = json.Unmarshal(val, &bootstrapAddrs)
	if err != nil {
		t.Fatal(err)
	}

	if len(bootstrapAddrs) == 0 {
		t.Errorf("expected number of boostrap addrs to be greater than 0 got %d",
			len(bootstrapAddrs))
	}

	// update bootstrap value
	err = cfg.SetKey("Bootstrap", []byte("[]"))
	if err != nil {
		t.Fatal(err)
	}

	// get bootstrap value again
	val, err = cfg.GetKey("Bootstrap")
	if err != nil {
		t.Fatal(err)
	}

	// check bootstrap value again
	err = json.Unmarshal(val, &bootstrapAddrs)
	if err != nil {
		t.Fatal(err)
	}

	if len(bootstrapAddrs) != 0 {
		t.Errorf("number of bootstrap addrs should be 0 but got: %d", len(bootstrapAddrs))
	}
}

func TestDefaultCofing(t *testing.T) {
	testCfg, err := NewDefaultConfig()
	if err != nil {
		t.Fatal(err)
	}

	val, err := testCfg.GetKey("Identity")
	if err != nil {
		t.Fatal(err)
	}

	var id *ipfs_config.Identity

	err = json.Unmarshal(val, &id)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(id.PeerID, "Qm") {
		t.Errorf("PeerID isn't prefixed by `Qm` got `%.2s` has prefix instead", id.PeerID)
	}
}
