package wallet

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
)

// DevicePubKey returns the device public key
func DevicePubKey() (crypto.PubKey, error) {
	// res, err := Instance.GetAll(token, wallet.Credential)
	// if err != nil {
	// 	return nil, err
	// }

	// for _, v := range res {
	// 	doc, err := did.ParseDocument(v)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	m, ok := did.LookupPublicKey(device_pub_key, doc)
	// 	if ok {
	// 		return crypto.UnmarshalPublicKey(m.Value)
	// 	}
	// }
	return nil, errors.New("no public key found")
}

// DevicePrivKey returns the device private key
func DevicePrivKey() (crypto.PrivKey, error) {
	// res, err := Instance.GetAll(token, wallet.Credential)
	// if err != nil {
	// 	return nil, err
	// }
	// for _, v := range res {
	// 	doc, err := did.ParseDocument(v)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	m, ok := did.LookupPublicKey(device_priv_key, doc)
	// 	if ok {
	// 		return crypto.UnmarshalPrivateKey(m.Value)
	// 	}
	// }
	return nil, errors.New("no private key found")
}
