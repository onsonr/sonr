package core

import (
	"encoding/json"
	"testing"
)

func testIDRequest(t *testing.T, node *Node, raw_json []byte) {
	id := struct {
		PeerID string `json:"id"`
	}{}

	err := json.Unmarshal(raw_json, &id)
	if err != nil {
		t.Fatal(err)
	}

	nodeID := node.ipfsMobile.IpfsNode.Identity.String()
	if nodeID != id.PeerID {
		t.Fatalf("PeerID should be equal to `%s` got `%s`", nodeID, id.PeerID)
	}
}

func testConfigRequest(t *testing.T, node *Node, raw_json []byte) {
	nodeConfig, err := node.ipfsMobile.IpfsNode.Repo.Config()
	if err != nil {
		t.Fatalf("unable to get node config: %s", err)
	}

	config := struct {
		Value string `json:"Value"`
	}{}

	err = json.Unmarshal(raw_json, &config)
	if err != nil {
		t.Fatal(err)
	}

	if nodeConfig.Identity.PeerID != config.Value {
		t.Fatalf("config.Identity should be equal to `%s` got `%s`", nodeConfig.Identity.PeerID, config.Value)
	}
}

func TestShell(t *testing.T) {
	sm, clean := testingSockmanager(t)
	defer clean()

	sockA, err := sm.NewSockPath()
	if err != nil {
		t.Fatal(err)
	}

	path, clean := testingTempDir(t, "repo")
	defer clean()

	node, clean := testingNode(t, path)
	defer clean()

	/// table cases
	// clients
	casesClient := map[string]struct{ MAddr string }{
		"tcp shell": {"/ip4/127.0.0.1/tcp/0"},
		"uds shell": {"/unix/" + sockA},
	}

	// commands
	casesCommand := []struct {
		Command      string
		Args         []string
		AssertMethod func(t *testing.T, node *Node, raw_json []byte)
	}{
		{"id", []string{}, testIDRequest},
		{"config", []string{"Identity.PeerID"}, testConfigRequest},
	}

	for clientk, clienttc := range casesClient {
		t.Run(clientk, func(t *testing.T) {
			maddr, err := node.ServeMultiaddr(clienttc.MAddr)
			if err != nil {
				t.Fatal(err)
			}

			shell := NewShell(maddr)
			for _, cmdtc := range casesCommand {
				t.Run(cmdtc.Command, func(t *testing.T) {
					req := shell.NewRequest(cmdtc.Command)
					for _, arg := range cmdtc.Args {
						req.Argument(arg)
					}

					res, err := req.Send()
					if err != nil {
						t.Error(err)
					}

					cmdtc.AssertMethod(t, node, res)
				})
			}

		})
	}
}
