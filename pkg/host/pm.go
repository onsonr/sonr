package host

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
  
	// "github.com/sonr-io/alice/example/utils"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type PeerManager struct {
	id       string
	host     SonrHost
	protocol protocol.ID
	peers    map[string]string
}

func NewPeerManager(id string, host SonrHost, protocol protocol.ID) *PeerManager {
	return &PeerManager{
		id:       id,
		host:     host,
		protocol: protocol,
		peers:    make(map[string]string),
	}
}

func (p *PeerManager) NumPeers() uint32 {
	return uint32(len(p.peers))
}

func (p *PeerManager) SelfID() string {
	return p.id
}

func (p *PeerManager) PeerIDs() []string {
	ids := make([]string, len(p.peers))
	i := 0
	for id := range p.peers {
		ids[i] = id
		i++
	}
	return ids
}

func (p *PeerManager) MustSend(peerID string, message interface{}) {
	p.host.Send(peer.ID(peerID), p.protocol, message.(protoreflect.ProtoMessage))
}

// EnsureAllConnected connects the host to specified peer and sends the message to it.
func (p *PeerManager) EnsureAllConnected() {
	var wg sync.WaitGroup

	for _, peerAddr := range p.peers {
		wg.Add(1)
		go connectToPeer(p.host, peerAddr, &wg)
	}
	wg.Wait()
}

// // AddPeers adds peers to peer list.
// func (p *PeerManager) AddPeers(peerPorts []int64) error {
// 	for _, peerPort := range peerPorts {
// 		peerID := utils.GetPeerIDFromPort(peerPort)
// 		peerAddr, err := getPeerAddr(peerPort)
// 		if err != nil {
// 			return err
// 		}
// 		p.peers[peerID] = peerAddr
// 	}
// 	return nil
// }

// generateIdentity generates a fixed key pair by using port as random source.
func generateIdentity(port int64) (crypto.PrivKey, error) {
	// Use the port as the randomness source in this example.
	// #nosec: G404: Use of weak random number generator (math/rand instead of crypto/rand)
	r := rand.New(rand.NewSource(port))

	// Generate a key pair for this host.
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.ECDSA, 2048, r)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

// getPeerAddr gets peer full address from port.
func getPeerAddr(port int64) (string, error) {
	priv, err := generateIdentity(port)
	if err != nil {
		return "", err
	}

	pid, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", port, pid), nil
}

func connectToPeer(host SonrHost, peerAddr string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ma, err := multiaddr.NewMultiaddr(peerAddr)
		if err != nil {
			panic(err)
		}

		pi, err := peer.AddrInfoFromP2pAddr(ma)
		if err != nil {
			panic(err)
		}
		// Connect the host to the peer.
		err = host.Connect(*pi)
		if err != nil {

			time.Sleep(3 * time.Second)
			continue
		}

		return
	}
}
