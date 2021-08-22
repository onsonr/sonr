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
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

type Account interface {
	// Config
	APIKeys() *md.APIKeys
	FilePath() string
	HandleSetPeer(peer *md.Peer, isPrimary bool)
	Save() error
	SetConnection(cr *md.ConnectionRequest)
	JoinNetwork(h sh.HostNode) *md.SonrError

	// Information
	Member() *md.Member
	DeviceID() string
	FirstName() string
	LastName() string
	Profile() *md.Profile
	UpdateContact(contact *md.Contact)

	// Keychain Management
	AccountKeys() *md.KeyPair
	CurrentDevice() *md.Device
	CurrentDeviceKeys() *md.KeyPair
	DeviceKeys() *md.KeyPair
	DevicePubKey() *md.KeyPair_Public
	GroupKeys() *md.KeyPair
	KeyChain() *md.KeyChain
	SignLinkPacket(resp *md.LinkResponse) *md.LinkPacket
	SignAuth(req *md.AuthRequest) *md.AuthResponse
	VerifyDevicePubKey(pub crypto.PubKey) bool
	VerifyGroupPubKey(pub crypto.PubKey) bool
	VerifyRead() *md.VerifyResponse

	// Linker Stream
	HandleLinkPacket(packet *md.LinkPacket)
	ReadFromLink(stream network.Stream)
	WriteToLink(stream network.Stream, resp *md.LinkResponse)

	// Room Management
	NewDefaultUpdateEvent(room *md.Room, id peer.ID) *md.RoomEvent
	NewUpdateEvent(room *md.Room, id peer.ID) *md.RoomEvent
	NewExitEvent(room *md.Room, id peer.ID) *md.RoomEvent
}

type userLinker struct {
	// General
	Account
	currentDevice *md.Device
	user          *md.User
	protocol      protocol.ID

	// Networking
	ctx          context.Context
	host         sh.HostNode
	topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler

	// Sync
	service       *DeviceService
	syncEvents    chan *md.SyncEvent
	activeDevices map[peer.ID]*md.Device
	room          *md.Room
	lastUpdated   int32
}

// Start Begins Account Manager
func OpenAccount(ir *md.InitializeRequest, d *md.Device) (Account, *md.SonrError) {
	// Fetch Key Pair
	keychain, err := d.Initialize(ir)
	if err != nil {
		return nil, err
	}

	// Check for existing account
	if d.GetFileSystem().GetSupport().IsFile(util.ACCOUNT_FILE) {
		// Load Account
		buf, err := d.GetFileSystem().GetSupport().ReadFile(util.ACCOUNT_FILE)
		if err != nil {
			return nil, err
		}

		// Unmarshal Account
		loadedAccount := &md.User{}
		serr := proto.Unmarshal(buf, loadedAccount)
		if serr != nil {
			return nil, md.NewError(serr, md.ErrorEvent_ACCOUNT_LOAD)
		}

		md.LogInfo(fmt.Sprintf("LoadedAccount: %s", loadedAccount.String()))

		// Set Account
		loadedAccount.KeyChain = keychain
		loadedAccount.Current = d
		loadedAccount.ApiKeys = ir.GetApiKeys()
		loadedAccount.State = ir.UserState()

		// Create Account Linker
		linker := &userLinker{
			user:          loadedAccount,
			currentDevice: ir.GetDevice(),
			room:          loadedAccount.NewDeviceRoom(),
		}
		return linker, nil
	} else {
		// Return User
		u := &md.User{
			KeyChain: keychain,
			Current:  d,
			ApiKeys:  ir.GetApiKeys(),
			State:    ir.UserState(),
			Devices:  make([]*md.Device, 0),
			Member: &md.Member{
				Reach:      md.Member_ONLINE,
				Associated: make([]*md.Peer, 0),
			},
		}

		// Create Account Linker
		linker := &userLinker{
			user:          u,
			currentDevice: ir.GetDevice(),
			room:          u.NewDeviceRoom(),
		}
		err := linker.Save()
		if err != nil {
			return nil, md.NewError(err, md.ErrorEvent_ACCOUNT_SAVE)
		}
		return linker, nil
	}
}

func (al *userLinker) JoinNetwork(h sh.HostNode) *md.SonrError {
	// Set host and context
	al.host = h
	al.ctx = context.Background()

	// Join Room
	topic, err := h.Pubsub().Join(al.room.GetName())
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_JOIN)
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_SUB)
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return md.NewError(err, md.ErrorEvent_ROOM_HANDLER)
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

// Update Account after Device Peer set for Member
func (al *userLinker) HandleSetPeer(p *md.Peer, isPrimary bool) {
	u := al.user
	if isPrimary {
		u.Member.Active = p
	} else {
		u.Member.Associated = append(u.Member.Associated, p)
	}
	al.Save()
}

// HandleLinkPacket Updates Account after LinkPacket is received
func (al *userLinker) HandleLinkPacket(lp *md.LinkPacket) {
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
			md.LogError(err)
			return
		}

		// Unmarshal linkPacket
		lp := &md.LinkPacket{}
		err = proto.Unmarshal(buf, lp)
		if err != nil {
			md.LogError(err)
			return
		}

		// Callback on Linked
		al.HandleLinkPacket(lp)
		stream.Close()
	}(msg.NewReader(stream), stream)
}

// SignAuth Method Signs Data with KeyPair
func (al *userLinker) SignAuth(req *md.AuthRequest) *md.AuthResponse {
	u := al.user
	// Create Prefix
	prefixResult := u.GetKeyChain().GetAccount().Sign(fmt.Sprintf("%s%s", req.GetSName(), al.DeviceID()))

	// Get Prefix Appended and Place
	prefix := util.Substring(prefixResult, 0, 16)

	// Get FingerPrint from Mnemonic and Place
	fingerprint := u.GetKeyChain().GetAccount().Sign(req.GetMnemonic())
	pubKey := u.GetKeyChain().GetAccount().PubKeyBase64()

	// Return Response
	return &md.AuthResponse{
		SignedPrefix:      prefix,
		SignedFingerprint: fingerprint,
		PublicKey:         pubKey,
		GivenSName:        req.GetSName(),
		GivenMnemonic:     req.GetMnemonic(),
	}
}

// UpdateContact Method Updates User Contact
func (al *userLinker) UpdateContact(c *md.Contact) {
	u := al.user
	u.Contact = c
	u.GetMember().UpdateProfile(c)
	u.Member.UpdateProfile(c)
	al.Save()
}

// WriteToLink writes Buffer onto from Stream to New Associated Device
func (al *userLinker) WriteToLink(stream network.Stream, resp *md.LinkResponse) {
	// Create Link Packet and Send
	linkPacket := &md.LinkPacket{
		Primary:   al.user.GetPrimary(),
		Secondary: resp.GetDevice(),
		KeyChain:  al.ExportKeychain(),
	}

	// Marshal linkPacket
	buf, err := proto.Marshal(linkPacket)
	if err != nil {
		md.LogError(err)
		return
	}

	// Concurrent Function
	go func(ws msg.WriteCloser) {
		err := ws.WriteMsg(buf)
		if err != nil {
			md.LogError(err)
		}
	}(msg.NewWriter(stream))
}
