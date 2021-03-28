package user

import (
	"hash/fnv"
	"log"
	"math"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Get Peer returns Users Contact ^ //
func (u *User) Contact() *md.Contact {
	return u.contact
}

// ^ Get Peer returns Users Current device ^ //
func (u *User) Device() *md.Device {
	return u.device
}

// ^ Get Peer returns Current Peer ^ //
func (u *User) Peer() *md.Peer {
	return u.peer
}

// ^ Peer returns Current Peer Info as Buffer
func (u *User) PeerBuf() []byte {
	// Convert to bytes
	buf, err := proto.Marshal(u.peer)
	if err != nil {
		log.Println(err)
		return nil
	}
	return buf
}

// ^ Get Peer returns Peers Profile ^ //
func (u *User) Profile() *md.Profile {
	return u.profile
}

// ^ Get Key: Returns Private key from disk if found ^ //
func (u *User) PrivateKey() (crypto.PrivKey, error) {
	// @ Get Private Key
	if ok := u.FS.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, err := u.FS.ReadFile(K_SONR_PRIV_KEY)
		if err != nil {
			return nil, err
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshalling identity private key")
		}

		// Set Key Ref
		return key, nil
	} else {
		// Create New Key
		privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return nil, errors.Wrap(err, "generating identity private key")
		}

		// Marshal Data
		buf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, errors.Wrap(err, "marshalling identity private key")
		}

		// Write Key to File
		_, err = u.FS.WriteFile(K_SONR_PRIV_KEY, buf)
		if err != nil {
			return nil, err
		}

		// Set Key Ref
		return privKey, nil
	}

}

// ^ Updates Current Contact Card ^
func (u *User) SetContact(newContact *md.Contact) error {
	// Set Node Contact
	u.contact = newContact

	// Update Peer Profile
	u.peer.Profile = &md.Profile{
		FirstName: newContact.GetFirstName(),
		LastName:  newContact.GetLastName(),
		Picture:   newContact.GetPicture(),
	}

	// Load User
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	// Save User
	if err := u.SaveUser(user); err != nil {
		return err
	}
	return nil
}

// ^ SetPeer configures Peer Ref from Host ID Reference ^ //
func (u *User) SetPeer(id peer.ID) error {
	// Initialize
	deviceID := u.device.GetId()

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(u.profile.GetUsername()))
	if err != nil {
		return err
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
	u.peer = &md.Peer{
		Id: &md.Peer_ID{
			Peer:   id.String(),
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  u.profile,
		Platform: u.device.Platform,
		Model:    u.device.Model,
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (u *User) SetPosition(facing float64, heading float64) {
	// Update User Values
	var faceDir float64
	var faceAnpd float64
	var headDir float64
	var headAnpd float64
	faceDir = math.Round(facing*100) / 100
	headDir = math.Round(heading*100) / 100
	desg := int((facing / 11.25) + 0.25)

	// Find Antipodal
	if facing > 180 {
		faceAnpd = math.Round((facing-180)*100) / 100
	} else {
		faceAnpd = math.Round((facing+180)*100) / 100
	}

	// Find Antipodal
	if heading > 180 {
		headAnpd = math.Round((heading-180)*100) / 100
	} else {
		headAnpd = math.Round((heading+180)*100) / 100
	}

	// Set Position
	u.peer.Position = &md.Position{
		Facing:           faceDir,
		FacingAntipodal:  faceAnpd,
		Heading:          headDir,
		HeadingAntipodal: headAnpd,
		Designation:      md.Position_Designation(desg % 32),
	}
}
