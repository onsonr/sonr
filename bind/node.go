package bind

import (
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/network"
	"github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Return URL Metadata, Helper Method
func GetURLMetadata(url string) []byte {
	// Get Link Data
	data, err := md.GetPageInfoFromUrl(url)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Failed to Parse URL"))
		return nil
	}

	// Marshal
	bytes, err := proto.Marshal(data)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Failed to Parse URL"))
		return nil
	}
	return bytes
}

// @ Join Existing Group
func (mn *Node) CreateRemote() []byte {
	if mn.isReady() {
		// Generate Word List
		_, wordList, err := network.RandomWords("english", 3)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Generating Random Quotes"))
			return nil
		}
		// Create Remote Request and Join Lobby
		remote := md.GetRemoteInfo(wordList)

		// Join Lobby
		tm, err := mn.client.JoinLobby(remote.Topic, true)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Joining Remote Topic"))
			return nil
		}

		// Set Topic
		mn.topics[remote.Topic] = tm

		// Marshal
		data, err := proto.Marshal(&remote)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Marshalling Remote Data"))
			return nil
		}
		return data
	}
	return nil
}

// @ Join Existing Group
func (mn *Node) JoinRemote(data []byte) {
	if mn.isReady() {
		// Unpackage Data
		remote := md.RemoteInfo{}
		err := proto.Unmarshal(data, &remote)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Unmarshalling Remote Data"))
			mn.error(err, "JoinRemote")
			return
		}

		// Join Lobby
		tm, err := mn.client.JoinLobby(remote.Topic, false)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Joining Remote Lobby"))
			mn.error(err, "JoinRemote")
			return
		}

		// Set Topic
		mn.topics[remote.Topic] = tm
	}
}

// @ Leave Existing Group
func (mn *Node) LeaveRemote(data []byte) {
	if mn.isReady() {
		// Unpackage Data
		remote := md.RemoteInfo{}
		err := proto.Unmarshal(data, &remote)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Unmarshalling Remote Info"))
			mn.error(err, "LeaveRemote")
			return
		}

		// Join Lobby
		err = mn.client.LeaveLobby(mn.topics[remote.Topic])
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Leaving Remote Lobby"))
			mn.error(err, "LeaveRemote")
			return
		}
	}
}

// @ Update proximity/direction and Notify Lobby
func (mn *Node) Update(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		udpate := &md.UpdateRequest{}
		if err := proto.Unmarshal(data, udpate); err != nil {
			sentry.CaptureException(errors.Wrap(err, "Unmarshalling Update Request"))
			return
		}

		// Update Peer
		mn.client.Peer.Update(udpate)

		// Notify Local Lobby
		err := mn.client.Update(mn.local)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Updating Local Lobby"))
			return
		}
	}
}

// @ Send Direct Message to Peer in Lobby
func (mn *Node) Message(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		req := &md.MessageRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			sentry.CaptureException(errors.Wrap(err, "Unmarshalling Message Request"))
			return
		}

		// Run Node Action
		err := mn.client.Message(mn.local, req.Message, req.To)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "Sending Message Request to Lobby"))
			return
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *Node) Invite(data []byte) {
	if mn.isReady() {
		// Update Status
		mn.setStatus(md.Status_PENDING)

		// Initialize from Request
		req := &md.InviteRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			sentry.CaptureException(errors.Wrap(err, "Unmarshalling Invite Request"))
			return
		}

		// Retreive Invite Topic
		var topic *topic.TopicManager
		if req.IsRemote && req.Remote != nil {
			topic = mn.topics[req.Remote.Topic]
		} else {
			topic = mn.local
		}

		// @ 2. Check Transfer Type
		if req.Type == md.InviteRequest_Contact || req.Type == md.InviteRequest_FlatContact {
			err := mn.client.InviteContact(req, topic, mn.user.Contact())
			if err != nil {
				sentry.CaptureException(errors.Wrap(err, "Inviting with Contact"))
				return
			}
		} else if req.Type == md.InviteRequest_URL {
			err := mn.client.InviteLink(req, topic)
			if err != nil {
				sentry.CaptureException(errors.Wrap(err, "Inviting with Link"))
				return
			}
		} else {
			// Invite With file
			err := mn.client.InviteFile(req, topic, mn.user.FileSystem())
			if err != nil {
				sentry.CaptureException(errors.Wrap(err, "Inviting with File"))
				return
			}
		}
	}
}

// @ Respond to an Invite with Decision
func (mn *Node) Respond(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		req := &md.RespondRequest{}
		if err := proto.Unmarshal(data, req); err != nil {
			sentry.CaptureException(errors.Wrap(err, "Unmarshalling Respond Request"))
			return
		}

		// Retreive Invite Topic
		var topic *topic.TopicManager
		if req.IsRemote && req.Remote != nil {
			topic = mn.topics[req.Remote.Topic]
		} else {
			topic = mn.local
		}

		mn.client.Respond(req, topic, mn.user.FileSystem(), mn.user.Contact())
		// Update Status
		if req.Decision {
			mn.setStatus(md.Status_INPROGRESS)
		} else {
			mn.setStatus(md.Status_AVAILABLE)
		}
	}
}
