package fs

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/di-dao/sonr/internal/local"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
)

// Helper function to parse IPFS path
func ParsePath(p string) (path.Path, error) {
	return path.NewPath(p)
}

// FetchVaultPath returns the path to the vault with the given name
func FetchVaultPath(name string) string {
	return VaultsFolder.Join(name).Path()
}

// OpenURL is a Helper function which opens a url in windows,linux,osx default browser
func OpenURL(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Helper function to get IPFS client
func getIPFSClient() (*rpc.HttpApi, error) {
	return local.GetIPFSClient()
}
