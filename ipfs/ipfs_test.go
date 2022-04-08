package ipfs

import (
	"context"
	"fmt"
	"log"
	"testing"

	iface "github.com/ipfs/interface-go-ipfs-core"
)

var tempNode iface.CoreAPI
var tempCid string

func TestAddFileTemp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	node, err := SpawnEphemeral(ctx)
	tempNode = node
	if err != nil {
		t.Error(err)
	}
	if node == nil {
		t.Errorf("SpawnDefault(ctx) resulted in nil result")
		return
	}

	data := []byte("Hello World!!!")
	resp, err := UploadData(ctx, data, node)
	if err != nil {
		t.Errorf("UploadData([]byte, coreAPI) resulted in status %d", resp.Status)
		t.Error(err)
	}

	tempCid = resp.Cid
}

func TestDonwloadFileTemp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// node, err := SpawnEphemeral(ctx)
	// if err != nil {
	// 	t.Error(err)
	// }
	// if node == nil {
	// 	t.Errorf("SpawnEphemeral(ctx) resulted in nil result")
	// 	return
	// }
	node := tempNode

	bootstrapNodes := []string{
		// IPFS Bootstrapper nodes.
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",

		// IPFS Cluster Pinning nodes
		"/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		"/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		"/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		"/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		"/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		"/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		"/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
		"/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",

		// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
		// "/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		// "/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
	}
	go func() {
		err := ConnectToPeers(ctx, node, bootstrapNodes)
		if err != nil {
			log.Printf("failed connect to peers: %s", err)
		}
	}()

	exampleCIDStr := "QmUaoioqU7bxezBQZkUcgcSyokatMY71sxsALxQmRRrHrj"
	resp, err := DownloadData(ctx, exampleCIDStr, node)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(resp)

	// fmt.Printf("Fetching a file from the network with CID %s\n", exampleCIDStr)
	// outputPath := outputBasePath + exampleCIDStr
	// testCID := icorepath.New(exampleCIDStr)

	// rootNode, err := ipfs.Unixfs().Get(ctx, testCID)
	// if err != nil {
	// 	panic(fmt.Errorf("Could not get file with CID: %s", err))
	// }

	// err = files.WriteTo(rootNode, outputPath)
	// if err != nil {
	// 	panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
	// }

	// fmt.Printf("Wrote the file to %s\n", outputPath)
}

// func TestCreatePerm(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	node, err := SpawnDefault(ctx)
// 	nodePerm = node
// 	fmt.Println(err)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if node == nil {
// 		t.Errorf("SpawnDefault(ctx) resulted in nil result")
// 	}
// }

// func TestUploadPerm(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	data := []byte("Hello World!!!")
// 	resp, err := UploadData(ctx, data, nodePerm)
// 	if err != nil {
// 		t.Errorf("UploadData([]byte, coreAPI) resulted in status %d", resp.Status)
// 		t.Error(err)
// 	}
// 	permCID = resp.Cid
// }
