package lobby

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/internal/models"
	net "github.com/sonr-io/core/pkg/net"
	"google.golang.org/protobuf/proto"
)

type Group struct {
	ctx          context.Context
	data         *md.Group
	lobby        *Lobby
	name         string
	groupsv      *GroupService
	subscription *pubsub.Subscription
	topic        *pubsub.Topic
}

// ^ Creates Group from BIP Words
func (lob *Lobby) JoinGroup(group string) error {
	// Join the local pubsub Topic
	topic, err := lob.pubSub.Join(lob.router.Topic(net.SetIDForGroup(group)))
	if err != nil {
		return err
	}

	// Subscribe to local Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return err
	}

	// Create Group
	g := &Group{
		ctx:          lob.ctx,
		lobby:        lob,
		name:         group,
		topic:        topic,
		subscription: sub,
		data: &md.Group{
			Name:    group,
			Size:    1,
			Members: make(map[string]*md.Peer),
		},
	}

	// Create PeerService
	groupsvServer := gorpc.NewServer(lob.host, lob.router.Exchange(net.SetIDForGroup(group)))
	gsv := GroupService{
		updatePeer: lob.updatePeer,
		getUser:    lob.call.Peer,
	}

	// Register Service
	err = groupsvServer.Register(&gsv)
	if err != nil {
		return err
	}

	// Set Service
	g.groupsv = &gsv

	// Add to Lobby
	lob.data.Groups[group] = g.data
	return nil
}

// ^ handleMessages pulls messages from the pubsub topic and pushes them onto the Messages channel. ^
func (group *Group) handleEvents() {
	// @ Create Topic Handler
	topicHandler, err := group.topic.EventHandler()
	if err != nil {
		log.Println(err)
		return
	}

	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := topicHandler.NextPeerEvent(group.ctx)
		if err != nil {
			topicHandler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			err := group.Exchange(lobEvent.Peer)
			if err != nil {
				group.lobby.call.Error(err, "Group-Fetch")
			}
		}

		if lobEvent.Type == pubsub.PeerLeave {
			group.removeMember(lobEvent.Peer)
		}
		md.GetState().NeedsWait()
	}
}

// ****************** //
// ** GRPC Service ** //
// ****************** //
// ExchangeArgs is Peer protobuf
type GroupArgs struct {
	Data []byte
}

// ExchangeResponse is also Peer protobuf
type GroupResponse struct {
	Data []byte
}

// Service Struct
type GroupService struct {
	getUser    md.ReturnPeer
	updatePeer md.UpdatePeer
}

// ^ Calls Invite on Remote Peer ^ //
func (gs *GroupService) ExchangeWith(ctx context.Context, args ExchangeArgs, reply *GroupResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Data, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers
	gs.updatePeer(remotePeer)
	userPeer := gs.getUser()

	// Convert Protobuf to bytes
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (g *Group) Exchange(id peer.ID) error {
	// Get Peer Data
	userPeer := g.lobby.call.Peer()
	msgBytes, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Initialize RPC
	rpcClient := gorpc.NewClient(g.lobby.host, g.lobby.router.Exchange(net.SetIDForGroup(g.name)))
	var reply ExchangeResponse
	var args ExchangeArgs
	args.Data = msgBytes

	// Call to Peer
	err = rpcClient.Call(id, "GroupService", "ExchangeWith", args, &reply)
	if err != nil {
		return err
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Data, remotePeer)

	// Send Error
	if err != nil {
		return err
	}

	// Update Peer with new data
	g.updateMember(remotePeer)
	return nil
}

// ^ removeMember removes Peer from Group ^
func (g *Group) removeMember(id peer.ID) {
	// Update Peer with new data
	delete(g.data.Members, id.String())
	g.data.Size = int32(len(g.data.Members)) + 1 // Account for User

	// Update in Lobby and Callback
	g.lobby.data.Groups[g.name] = g.data
	g.lobby.Refresh()
}

// ^ updateMember changes Peer values in Group ^
func (g *Group) updateMember(peer *md.Peer) {
	// Update Peer with new data
	g.data.Members[peer.Id.Peer] = peer
	g.data.Size = int32(len(g.data.Members)) + 1 // Account for User

	// Update in Lobby and Callback
	g.lobby.data.Groups[g.name] = g.data
	g.lobby.Refresh()
}
