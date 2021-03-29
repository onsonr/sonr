package peer

import (
	"hash/fnv"
	"log"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type PeerNode struct {
	device  *md.Device
	peer    *md.Peer
	profile *md.Profile
}

// ^ Creates New Peer Node ^ //
func NewPeer(cr *md.ConnectionRequest, id peer.ID) (*PeerNode, error) {
	// Initialize
	deviceID := cr.Device.GetId()
	profile := &md.Profile{
		Username:  cr.GetUsername(),
		FirstName: cr.Contact.GetFirstName(),
		LastName:  cr.Contact.GetLastName(),
		Picture:   cr.Contact.GetPicture(),
		Platform:  cr.Device.GetPlatform(),
	}

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.GetUsername()))
	if err != nil {
		return nil, err
	}

	// Check if ID not provided
	if deviceID == "" {
		// Generate ID
		if id, err := mid.ProtectedID("Sonr"); err != nil {
			log.Println(err)
			deviceID = ""
		} else {
			deviceID = id
		}
	}

	// Set Peer
	return &PeerNode{
		device:  cr.Device,
		profile: profile,
		peer: &md.Peer{
			Id: &md.Peer_ID{
				Peer:   id.String(),
				Device: deviceID,
				User:   userID.Sum32(),
			},
			Profile:  profile,
			Platform: cr.Device.Platform,
			Model:    cr.Device.Model,
		},
	}, nil
}

// ^ Returns Peer as Buffer ^ //
func (pn *PeerNode) Buffer() []byte {
	buf, err := proto.Marshal(pn.peer)
	if err != nil {
		return nil
	}
	return buf
}

// ^ Get Returns Peer Proto ^ //
func (pn *PeerNode) Get() *md.Peer {
	return pn.peer
}

// ^ Checks for Host Peer ID is Same ^ //
func (pn *PeerNode) IsPeerID(pid peer.ID) bool {
	return pn.peer.Id.Peer == pid.String()
}

// ^ Checks for Host Peer ID String is Same ^ //
func (pn *PeerNode) IsPeerIDString(pid string) bool {
	return pn.peer.Id.Peer == pid
}

// ^ Checks for Host Peer ID String is not Same ^ //
func (pn *PeerNode) IsNotPeerIDString(pid string) bool {
	return pn.peer.Id.Peer != pid
}
