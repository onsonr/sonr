package wallet

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	"github.com/hyperledger/aries-framework-go/pkg/wallet"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/device"
)

// createDefaultKeys creates the default keys
func createDefaultKeys(sname string) error {
	res, err := Instance.GetAll(token, wallet.Credential)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		doc, err := newDeviceDID()
		if err != nil {
			return err
		}

		raw, err := doc.MarshalJSON()
		if err != nil {
			return err
		}

		// convert doc to raw json message
		logger.Infof("Created Device DID: %v", string(raw))
		err = Instance.Add(token, wallet.Credential, raw)
		if err != nil {
			return err
		}
		logger.Infof("Created new Key")
	}
	return nil
}

// newDeviceDID returns the device DID
//
// {
//   "@context": ["https://w3id.org/wallet/v1"],
//   "id": "urn:uuid:e8fc7810-9524-11ea-bb37-0242ac130002",
//   "name": "My Test Key 2",
//   "image": "https://via.placeholder.com/150",
//   "description" : "For testing only, totally compromised.",
//   "tags": ["professional", "organization", "compromised"],
//   "correlation": ["4058a72a-9523-11ea-bb37-0242ac130002"],
//   "controller": "did:key:z6MkjjCpsoQrwnEmqHzLdxWowXk5gjbwor4urC1RPDmGeV8r",
//   "type": "Ed25519VerificationKey2018",
//   "privateKeyBase58": "3CQCBKF3Mf1tU5q1FLpHpbxYrNYxLiZk4adDtfyPEfc39Wk6gsTb2qoc1ZtpqzJYdM1rG4gpaD3ZVKdkiDrkLF1p",
//   "publicKeyBase58": "6GwnHZARcEkJio9dxPYy6SC5sAL6PxpZAB6VYwoFjGMU"
// }
func newDeviceDID() (*did.Doc, error) {
	privKey, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	privBuf, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	devid, err := device.ID()
	if err != nil {
		return nil, err
	}

	devicePubVerify := did.NewVerificationMethodFromBytes("did:sonr:device-public-key", pubKey.Type().String(), devid, pubBuf)
	devicePrivVerify := did.NewVerificationMethodFromBytes("did:sonr:device-private-key", privKey.Type().String(), devid, privBuf)
	verificationMethod := []did.VerificationMethod{*devicePubVerify, *devicePrivVerify}
	didDoc := did.BuildDoc(did.WithVerificationMethod(verificationMethod))
	didDoc.ID = fmt.Sprintf("did:sonr:%s", devid)
	return didDoc, nil
}

// DevicePubKey returns the device public key
func DevicePubKey() (crypto.PubKey, error) {
	res, err := Instance.GetAll(token, wallet.Credential)
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		doc, err := did.ParseDocument(v)
		if err != nil {
			return nil, err
		}
		m, ok := did.LookupPublicKey("did:sonr:device-public-key", doc)
		if ok {
			return crypto.UnmarshalPublicKey(m.Value)
		}
	}
	return nil, errors.New("no public key found")
}

// DevicePrivKey returns the device private key
func DevicePrivKey() (crypto.PrivKey, error) {
	res, err := Instance.GetAll(token, wallet.Credential)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		doc, err := did.ParseDocument(v)
		if err != nil {
			return nil, err
		}
		m, ok := did.LookupPublicKey("did:sonr:device-private-key", doc)
		if ok {
			return crypto.UnmarshalPrivateKey(m.Value)
		}
	}
	return nil, errors.New("no private key found")
}
