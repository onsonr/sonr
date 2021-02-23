package host

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/internal/models"
)

// ^ Bootstrap Nodes ^ //
type Config struct {
	P2P struct {
		RDVP []struct {
			Maddr string `json:"maddr"`
		} `json:"rdvp"`
		RelayHack []string `json:"relayHack" yaml:"relayHack"`
	} `json:"p2p"`
}

var config Config

func init() {
	input := `
{
  "p2p": {
    "rdvp": [
      {
        "maddr": "/ip4/51.159.21.214/tcp/4040/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p"
      },
      {
        "maddr": "/ip4/51.159.21.214/udp/4040/quic/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p"
      },
      {
        "maddr": "/ip4/51.15.25.224/tcp/4040/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ"
      },
      {
        "maddr": "/ip4/51.15.25.224/udp/4040/quic/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ"
      },
      {
        "maddr": "/ip4/51.75.127.200/tcp/4141/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
      },
      {
        "maddr": "/ip4/51.75.127.200/udp/4141/quic/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
      }
    ],
    "relayHack": [
      "/ip4/51.159.21.214/udp/4040/quic/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p",
      "/ip4/51.15.25.224/udp/4040/quic/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ",
      "/ip4/51.75.127.200/udp/4141/quic/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
    ]
  }
}`

	err := json.Unmarshal([]byte(input), &config)
	if err != nil {
		panic(err)
	}
}

// ^ Get Keys: Returns Private/Public keys from disk if found ^ //
func getKeys(dir *md.Directories) (crypto.PrivKey, error) {
	// Set Path
	path := filepath.Join(dir.Documents, ".sonr-priv-key")

	// @ Path Doesnt Exist Generate Keys
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Generate Keys
		privKey, _, err := crypto.GenerateRSAKeyPair(2048, rand.Reader)
		if err != nil {
			return nil, err
		}

		// Get Key Bytes
		privDat, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, err
		}

		// Write Private/Pub To File
		err = ioutil.WriteFile(path, privDat, 0644)
		if err != nil {
			return nil, err
		}
		return privKey, nil
	}
	// @ Keys Exist Load Keys
	// Load Private Key Bytes from File
	privDat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal PrivKey from Bytes
	privKey, err := crypto.UnmarshalPrivateKey(privDat)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}
