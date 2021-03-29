package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/internal/network"
	"github.com/sonr-io/core/pkg/topic"
	"google.golang.org/protobuf/proto"
)

// @ Return URL Metadata, Helper Method
func GetURLMetadata(url string) []byte {
	// Get Link Data
	data, err := md.GetPageInfoFromUrl(url)
	if err != nil {
		log.Println(err, " Failed to Parse URL")
	}

	// Marshal
	bytes, err := proto.Marshal(data)
	if err != nil {
		log.Println(err, " Failed to Parse URL")
	}
	return bytes
}

// @ Join Existing Group
func (mn *MobileNode) CreateRemote() []byte {
	if mn.isReady() {
		// Generate Word List
		_, wordList, err := network.RandomWords("english", 3)
		if err != nil {
			return nil
		}
		// Create Remote Request and Join Lobby
		remote := md.GetRemoteInfo(wordList)

		// Join Lobby
		tm, err := mn.node.JoinLobby(remote.Topic)
		if err != nil {
			mn.error(err, "JoinRemote")
			return nil
		}

		// Set Topic
		mn.topics[remote.Topic] = tm

		// Marshal
		data, err := proto.Marshal(&remote)
		if err != nil {
			return nil
		}
		return data
	}
	return nil
}

// @ Join Existing Group
func (mn *MobileNode) JoinRemote(data []byte) {
	if mn.isReady() {
		// Unpackage Data
		remote := md.RemoteInfo{}
		err := proto.Unmarshal(data, &remote)
		if err != nil {
			mn.error(err, "JoinRemote")
			return
		}

		// Join Lobby
		tm, err := mn.node.JoinLobby(remote.Topic)
		if err != nil {
			mn.error(err, "JoinRemote")
			return
		}

		// Set Topic
		mn.topics[remote.Topic] = tm
	}
}

// @ Update proximity/direction and Notify Lobby
func (mn *MobileNode) Update(facing float64, heading float64) {
	if mn.isReady() {
		mn.peer.SetPosition(facing, heading)
		err := mn.node.Update(mn.local)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// @ Send Direct Message to Peer in Lobby
func (mn *MobileNode) Message(msg string, to string) {
	if mn.isReady() {
		err := mn.node.Message(mn.local, msg, to)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *MobileNode) Invite(reqBytes []byte) {
	if mn.isReady() {
		// Update Status
		mn.setStatus(md.Status_PENDING)

		// Initialize from Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(reqBytes, req); err != nil {
			log.Println(err)
			return
		}

		// Retreive Invite Topic
		var topic *topic.TopicManager
		if req.IsRemote {
			topic = mn.topics[req.Topic]
		} else {
			topic = mn.local
		}

		// @ 2. Check Transfer Type
		if req.Type == md.InviteRequest_Contact {
			err := mn.node.InviteContact(req, topic, mn.user.Contact())
			if err != nil {
				log.Println(err)
				return
			}
		} else if req.Type == md.InviteRequest_URL {
			err := mn.node.InviteLink(req, topic)
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			// Invite With file
			err := mn.node.InviteFile(req, topic, mn.user.FS)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

// @ Respond to an Invite with Decision
func (mn *MobileNode) Respond(decs bool) {
	if mn.isReady() {
		mn.node.Respond(decs, mn.user.FS, mn.local, mn.user.Contact())
		// Update Status
		if decs {
			mn.setStatus(md.Status_INPROGRESS)
		} else {
			mn.setStatus(md.Status_AVAILABLE)
		}
	}
}

// ** User Actions ** //
// @ Updates Current Contact Card
func (mn *MobileNode) SetContact(conBytes []byte) {
	if mn.isReady() {
		// Unmarshal Data
		newContact := &md.Contact{}
		err := proto.Unmarshal(conBytes, newContact)
		if err != nil {
			log.Println(err)
			return
		}

		// Update Node Profile
		err = mn.user.SetContact(newContact)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
