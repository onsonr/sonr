package account

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
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
	PrepareToLink(resp *md.LinkResponse)
	HandleLinkPacket(packet *md.LinkPacket)
	ReadFromLink(stream network.Stream)
	WriteToLink(stream network.Stream)
}

type accountLinker struct {
	// General
	Account
	account  *md.User
	protocol protocol.ID

	// Networking
	ctx          context.Context
	host         sh.HostNode
	Topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler

	// Sync
	service       *DeviceService
	syncEvents    chan *md.SyncEvent
	activeDevices []*md.Device
	room          *md.Room
	lastUpdated   int32

	// Linker Stream
	linkPacket *md.LinkPacket
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
		linker := &accountLinker{
			account: loadedAccount,
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
		linker := &accountLinker{
			account: u,
		}
		err := linker.Save()
		if err != nil {
			return nil, md.NewError(err, md.ErrorEvent_ACCOUNT_SAVE)
		}
		return linker, nil
	}
}

func (al *accountLinker) JoinNetwork(h sh.HostNode) *md.SonrError {
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
	al.Topic = topic
	al.subscription = sub

	// Start Sync
	serr := al.initService()
	if err != nil {
		return serr
	}
	return nil
}

// Update Account after Device Peer set for Member
func (al *accountLinker) HandleSetPeer(p *md.Peer, isPrimary bool) {
	u := al.account
	if isPrimary {
		u.Member.Active = p
	} else {
		u.Member.Associated = append(u.Member.Associated, p)
	}
	al.Save()
}

// Update Account after LinkPacket is received
func (al *accountLinker) HandleLinkPacket(lp *md.LinkPacket) {
	u := al.account
	u.Primary = lp.Primary
	u.Devices = append(u.Devices, lp.Secondary)
	u.GetCurrent().ReplaceKeyChain(lp.GetKeyChain())
	al.Save()
}

// Node Device is Sending Primary Account Information
func (al *accountLinker) PrepareToLink(resp *md.LinkResponse) {
	u := al.account
	al.linkPacket = &md.LinkPacket{
		Primary:   u.GetPrimary(),
		Secondary: resp.GetDevice(),
		KeyChain:  al.ExportKeychain(),
	}
}

func (al *accountLinker) ReadFromLink(stream network.Stream) {
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

// Method Signs Data with KeyPair
func (al *accountLinker) SignAuth(req *md.AuthRequest) *md.AuthResponse {
	u := al.account
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

// Method Signs Packet with Keys
func (al *accountLinker) SignLinkPacket(resp *md.LinkResponse) *md.LinkPacket {
	u := al.account
	return &md.LinkPacket{
		Primary:   u.GetPrimary(),
		Secondary: resp.GetDevice(),
		KeyChain:  al.ExportKeychain(),
	}
}

// Method Updates User Contact
func (al *accountLinker) UpdateContact(c *md.Contact) {
	u := al.account
	u.Contact = c
	u.GetMember().UpdateProfile(c)
	al.Save()
}

// Method Verifies the Device Link Public Key
func (al *accountLinker) VerifyDevicePubKey(pub crypto.PubKey) bool {
	u := al.account
	return u.GetKeyChain().Device.VerifyPubKey(pub)
}

// Method Updates User Contact
func (al *accountLinker) VerifyRead() *md.VerifyResponse {
	u := al.account
	kp := u.GetKeyChain().GetAccount()
	return &md.VerifyResponse{
		PublicKey: kp.PubKeyBase64(),
		ShortID:   u.GetCurrent().ShortID(),
	}
}

// Write Buffer onto from Stream to New Associated Device
func (al *accountLinker) WriteToLink(stream network.Stream) {
	// Marshal linkPacket
	buf, err := proto.Marshal(al.linkPacket)
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
