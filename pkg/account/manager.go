package account

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	ps "github.com/libp2p/go-libp2p-pubsub"
	msg "github.com/libp2p/go-msgio"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

type Account interface {
	// Config
	APIKeys() *data.APIKeys
	FilePath() string
	Save() error
	JoinNetwork(h sh.HostNode, cr *data.ConnectionRequest, p *data.Peer) *data.SonrError

	// Information
	Member() *data.Member
	DeviceID() string
	FirstName() string
	LastName() string
	Profile() *data.Profile
	UpdateContact(contact *data.Contact)

	// Keychain Management
	AccountKeys() *data.KeyPair
	CurrentDevice() *data.Device
	CurrentDeviceKeys() *data.KeyPair
	DeviceKeys() *data.KeyPair
	DevicePubKey() *data.KeyPair_Public
	GroupKeys() *data.KeyPair
	KeyChain() *data.KeyChain
	SignLinkPacket(resp *data.LinkResponse) *data.LinkPacket
	SignAuth(req *data.AuthRequest) *data.AuthResponse
	SignInvite(i *data.InviteRequest) *data.InviteRequest
	VerifyDevicePubKey(pub crypto.PubKey) bool
	VerifyGroupPubKey(pub crypto.PubKey) bool
	VerifyRead() *data.VerifyResponse

	// Linker Stream
	HandleLinkPacket(packet *data.LinkPacket)
	ReadFromLink(stream network.Stream)
	WriteToLink(stream network.Stream, resp *data.LinkResponse)

	// Room Management
	NewDefaultUpdateEvent(room *data.Room, id peer.ID) *data.RoomEvent
	NewUpdateEvent(room *data.Room, id peer.ID) *data.RoomEvent
	NewExitEvent(room *data.Room, id peer.ID) *data.RoomEvent

	// Status Management
	SetAvailable(val bool) *data.StatusEvent
	SetConnected(val bool) *data.StatusEvent
	SetStatus(newStatus data.Status) *data.StatusEvent
}

type userLinker struct {
	// General
	Account
	user     *data.User
	protocol protocol.ID

	// Networking
	ctx          context.Context
	host         sh.HostNode
	topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler

	// Sync
	service       *DeviceService
	syncEvents    chan *data.SyncEvent
	activeDevices map[peer.ID]*data.Device
	room          *data.Room
	lastUpdated   int32
}

// Start Begins Account Manager
func OpenAccount(ir *data.InitializeRequest, d *data.Device) (Account, *data.SonrError) {
	// Fetch Key Pair
	keychain, err := d.Initialize(ir)
	if err != nil {
		return nil, err
	}

	// Check for existing account
	if d.GetFileSystem().GetSupport().IsFile(util.ACCOUNT_FILE) {
		linker, err := loadLinker(ir, d, keychain)
		if err != nil {
			data.LogError(err.Error)
			data.LogInfo("Failed to load account, creating new one...")
			linker, err := newLinker(ir, d, keychain)
			if err != nil {
				data.LogError(err.Error)
				return nil, err
			}
			return linker, nil
		}
		return linker, nil
	} else {
		linker, err := newLinker(ir, d, keychain)
		if err != nil {
			data.LogError(err.Error)
			return nil, err
		}
		return linker, nil
	}
}

// JoinNetwork Method Joins Network, Updates Profile, Sets Membership
func (al *userLinker) JoinNetwork(h sh.HostNode, cr *data.ConnectionRequest, p *data.Peer) *data.SonrError {
	// Initialize Networking Param
	al.ctx = context.Background()
	al.host = h

	// Initialize Account Params
	al.user.Current.Location = cr.GetLocation()
	al.user.PushToken = cr.GetPushToken()
	al.user.SName = cr.GetContact().GetProfile().GetSName()
	al.user.Contact = cr.GetContact()

	// Set Member
	al.user.Member.Active = p
	al.user.Member.UpdateProfile(cr.GetContact())
	al.Save()

	al.room = al.user.NewDeviceRoom()
	al.activeDevices = make(map[peer.ID]*data.Device, 0)
	al.syncEvents = make(chan *data.SyncEvent)

	// Join Room
	topic, err := al.host.Pubsub().Join(fmt.Sprintf("/sonr/device/%s", al.user.GetSName()))
	if err != nil {
		return data.NewError(err, data.ErrorEvent_ROOM_JOIN)
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		return data.NewError(err, data.ErrorEvent_ROOM_SUB)
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return data.NewError(err, data.ErrorEvent_ROOM_HANDLER)
	}

	// Create Lobby Manager
	al.eventHandler = handler
	al.topic = topic
	al.subscription = sub

	// Start Sync
	serr := al.initService()
	if err != nil {
		return serr
	}
	return nil
}

// HandleLinkPacket Updates Account after LinkPacket is received
func (al *userLinker) HandleLinkPacket(lp *data.LinkPacket) {
	u := al.user
	u.Primary = lp.Primary
	u.Devices = append(u.Devices, lp.Secondary)
	u.GetCurrent().ReplaceKeyChain(lp.GetKeyChain())
	al.Save()
}

// ReadFromLink reads stream for link packet
func (al *userLinker) ReadFromLink(stream network.Stream) {
	// Concurrent Function
	go func(rs msg.ReadCloser, stream network.Stream) {
		buf, err := rs.ReadMsg()
		if err != nil {
			data.LogError(err)
			return
		}

		// Unmarshal linkPacket
		lp := &data.LinkPacket{}
		err = proto.Unmarshal(buf, lp)
		if err != nil {
			data.LogError(err)
			return
		}

		// Callback on Linked
		al.HandleLinkPacket(lp)
		stream.Close()
	}(msg.NewReader(stream), stream)
}

// SignAuth Method Signs Data with KeyPair
func (al *userLinker) SignAuth(req *data.AuthRequest) *data.AuthResponse {
	u := al.user
	// Create Prefix
	prefixResult := u.GetKeyChain().GetAccount().Sign(fmt.Sprintf("%s%s", req.GetSName(), al.DeviceID()))

	// Get Prefix Appended and Place
	prefix := util.Substring(prefixResult, 0, 16)

	// Get FingerPrint from Mnemonic and Place
	fingerprint := u.GetKeyChain().GetAccount().Sign(req.GetMnemonic())
	pubKey := u.GetKeyChain().GetAccount().PubKeyBase64()

	// Return Response
	return &data.AuthResponse{
		SignedPrefix:      prefix,
		SignedFingerprint: fingerprint,
		PublicKey:         pubKey,
		GivenSName:        req.GetSName(),
		GivenMnemonic:     req.GetMnemonic(),
	}
}

// UpdateContact Method Updates User Contact
func (al *userLinker) UpdateContact(c *data.Contact) {
	al.user.Contact = c
	al.user.GetMember().UpdateProfile(c)
	al.user.Member.UpdateProfile(c)
	al.Save()
}

// WriteToLink writes Buffer onto from Stream to New Associated Device
func (al *userLinker) WriteToLink(stream network.Stream, resp *data.LinkResponse) {
	// Create Link Packet and Send
	linkPacket := &data.LinkPacket{
		Primary:   al.user.GetPrimary(),
		Secondary: resp.GetDevice(),
		KeyChain:  al.ExportKeychain(),
	}

	// Marshal linkPacket
	buf, err := proto.Marshal(linkPacket)
	if err != nil {
		data.LogError(err)
		return
	}

	// Concurrent Function
	go func(ws msg.WriteCloser) {
		err := ws.WriteMsg(buf)
		if err != nil {
			data.LogError(err)
		}
	}(msg.NewWriter(stream))
}
