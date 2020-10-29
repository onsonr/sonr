package core

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func TestNode(t *testing.T) {
	path, clean := testingTempDir(t, "repo")
	defer clean()

	repo, clean := testingRepo(t, path)
	defer clean()

	node, err := NewNode(repo)
	if err != nil {
		t.Fatal(err)
	}

	if err := node.Close(); err != nil {
		t.Error(err)
	}
}
func TestNodeServeAPI(t *testing.T) {
	t.Run("tpc api", func(t *testing.T) {
		path, clean := testingTempDir(t, "tpc_repo")
		defer clean()

		node, clean := testingNode(t, path)
		defer clean()

		smaddr, err := node.ServeTCPAPI("0")
		if err != nil {
			t.Fatal(err)
		}

		maddr, err := ma.NewMultiaddr(smaddr)
		if err != nil {
			t.Fatal(err)
		}

		addr, err := manet.ToNetAddr(maddr)
		if err != nil {
			t.Fatal(err)
		}

		url := fmt.Sprintf("http://%s/api/v0/id", addr.String())
		client := http.Client{Timeout: 5 * time.Second}

		_, err = client.Get(url)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("uds api", func(t *testing.T) {
		path, clean := testingTempDir(t, "uds_repo")
		defer clean()

		sockdir, clean := testingTempDir(t, "uds_api")
		defer clean()

		node, clean := testingNode(t, path)
		defer clean()

		sock := filepath.Join(sockdir, "sock")

		err := node.ServeUnixSocketAPI(sock)
		if err != nil {
			t.Fatal(err)
		}

		client := http.Client{
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", sock)
				},
			},
		}

		_, err = client.Get("http://unix/api/v0/id")
		if err != nil {
			t.Fatal(err)
		}
	})
}
