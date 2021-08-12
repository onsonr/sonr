package models

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	util "github.com/sonr-io/core/pkg/util"
)

// Method Initializes User Info Struct ^ //
func InitAccount(ir *InitializeRequest, d *Device) (*Account, *SonrError) {
	// Fetch Key Pair
	keychain, err := d.Initialize(ir)
	if err != nil {
		return nil, err
	}

	// Return User
	u := &Account{
		KeyChain: keychain,
		Primary:  d,
		ApiKeys:  ir.GetApiKeys(),
		State:    ir.AccountState(),
		Devices:  make([]*Device, 0),
		Member: &Member{
			Reach:      Member_ONLINE,
			Associated: make([]*Peer, 0),
		},
	}
	return u, nil
}

// Set the User with ConnectionRequest
func (u *Account) SetConnection(cr *ConnectionRequest) {
	// Initialize Params
	u.PushToken = cr.GetPushToken()
	u.SName = cr.GetContact().GetProfile().GetSName()
	u.Contact = cr.GetContact()
	u.Member.PushToken = cr.GetPushToken()
}

// Update Account after Device Peer set for Member
func (u *Account) HandleSetPeer(p *Peer, isPrimary bool) {
	if isPrimary {
		u.Member.Primary = p
	} else {
		u.Member.Associated = append(u.Member.Associated, p)
	}
}

// Checks Whether User is Ready to Communicate
func (u *Device) IsReady() bool {
	return u.Contact != nil && u.Location != nil && u.Status != Status_DEFAULT
}

// Return Client API Keys
func (u *Account) APIKeys() *APIKeys {
	return u.GetApiKeys()
}

// Method Returns DeviceID
func (u *Account) DeviceID() string {
	return u.GetPrimary().GetId()
}

// Method Returns Profile First Name
func (u *Account) FirstName() string {
	return u.GetContact().GetProfile().GetFirstName()
}

// Method Returns Profile Last Name
func (u *Account) LastName() string {
	return u.GetContact().GetProfile().GetLastName()
}

// Method Returns Profile
func (u *Account) Profile() *Profile {
	return u.GetContact().GetProfile()
}

// Method Returns Account KeyPair
func (u *Account) AccountKeys() *KeyPair {
	if u.Primary.HasAccountKeys() {
		return u.GetKeyChain().GetAccount()
	}
	return nil
}

// Method Returns Device Link Public Key
func (u *Account) DevicePubKey() *KeyPair_Public {
	if u.Primary.HasDeviceKeys() {
		return u.GetKeyChain().GetDevice().GetPublic()
	}
	return nil
}

// Method Signs Data with KeyPair
func (u *Account) Sign(req *AuthRequest) *AuthResponse {
	// Create Prefix
	prefixResult := u.GetKeyChain().GetAccount().Sign(fmt.Sprintf("%s%s", req.GetSName(), u.DeviceID()))

	// Get Prefix Appended and Place
	prefix := util.Substring(prefixResult, 0, 16)

	// Get FingerPrint from Mnemonic and Place
	fingerprint := u.GetKeyChain().GetAccount().Sign(req.GetMnemonic())
	pubKey := u.GetKeyChain().GetAccount().PubKeyBase64()

	// Return Response
	return &AuthResponse{
		SignedPrefix:      prefix,
		SignedFingerprint: fingerprint,
		PublicKey:         pubKey,
		GivenSName:        req.GetSName(),
		GivenMnemonic:     req.GetMnemonic(),
	}
}

// Method Updates User Contact
func (u *Account) UpdateContact(c *Contact) {
	u.Contact = c
	u.GetMember().UpdateProfile(c)
}

// Method Verifies the Device Link Public Key
func (u *Account) VerifyDevicePubKey(pub crypto.PubKey) bool {
	return u.GetKeyChain().Device.VerifyPubKey(pub)
}

// Method Updates User Contact
func (u *Account) VerifyRead() *VerifyResponse {
	kp := u.GetKeyChain().GetAccount()
	return &VerifyResponse{
		PublicKey: kp.PubKeyBase64(),
		ShortID:   u.GetPrimary().ShortID(),
	}
}
