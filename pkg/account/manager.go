package account

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	ps "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

type AccountManager interface {
	// Config
	APIKeys() *md.APIKeys
	FilePath() string
	HandleSetPeer(peer *md.Peer, isPrimary bool)
	Save() error
	SetConnection(cr *md.ConnectionRequest)
	JoinDeviceRoom(h sh.HostNode) *md.SonrError

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
	HandleLinkPacket(packet *md.LinkPacket)
	ReadFromLink(stream network.Stream)
	WriteToLink(stream network.Stream)
}

type accountLinker struct {
	// General
	AccountManager
	account  *md.Account
	protocol protocol.ID

	// Networking
	ctx          context.Context
	host         sh.HostNode
	Topic        *ps.Topic
	subscription *ps.Subscription
	eventHandler *ps.TopicEventHandler

	// Sync
	verify        *DeviceService
	syncEvents    chan *md.SyncEvent
	activeDevices []*md.Device
	room          *md.Room
}

// Start Begins Account Manager
func StartAccount(ir *md.InitializeRequest, d *md.Device) (AccountManager, *md.SonrError) {
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
		loadedAccount := &md.Account{}
		serr := proto.Unmarshal(buf, loadedAccount)
		if serr != nil {
			return nil, md.NewError(serr, md.ErrorEvent_ACCOUNT_LOAD)
		}

		md.LogInfo(fmt.Sprintf("LoadedAccount: %s", loadedAccount.String()))

		// Set Account
		loadedAccount.KeyChain = keychain
		loadedAccount.Current = d
		loadedAccount.ApiKeys = ir.GetApiKeys()
		loadedAccount.State = ir.AccountState()

		// Create Account Linker
		linker := &accountLinker{
			account: loadedAccount,
		}
		return linker, nil
	} else {
		// Return User
		u := &md.Account{
			KeyChain: keychain,
			Current:  d,
			ApiKeys:  ir.GetApiKeys(),
			State:    ir.AccountState(),
			Devices:  make([]*md.Device, 0),
			Member: &md.Member{
				Reach:      md.Member_ONLINE,
				Associated: make([]*md.Peer, 0),
			},
		}
		err := u.Save()
		if err != nil {
			return nil, md.NewError(err, md.ErrorEvent_ACCOUNT_SAVE)
		}

		// Create Account Linker
		linker := &accountLinker{
			account: u,
		}
		return linker, nil
	}
}

func (al *accountLinker) JoinDeviceRoom(h sh.HostNode) *md.SonrError {
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
	u.Save()
}

// Update Account after LinkPacket is received
func (al *accountLinker) HandleLinkPacket(lp *md.LinkPacket) {
	u := al.account
	u.Primary = lp.Primary
	u.Devices = append(u.Devices, lp.Secondary)
	u.GetCurrent().ReplaceKeyChain(lp.GetKeyChain())
	u.Save()
}

// Method Signs Data with KeyPair
func (al *accountLinker) SignAuth(req *md.AuthRequest) *md.AuthResponse {
	u := al.account
	// Create Prefix
	prefixResult := u.GetKeyChain().GetAccount().Sign(fmt.Sprintf("%s%s", req.GetSName(), u.DeviceID()))

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
		KeyChain:  u.ExportKeychain(),
	}
}

// Method Updates User Contact
func (al *accountLinker) UpdateContact(c *md.Contact) {
	u := al.account
	u.Contact = c
	u.GetMember().UpdateProfile(c)
	u.Save()
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
