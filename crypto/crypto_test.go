package crypto

import (
	"bufio"
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	crypto_rand "crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/okx/go-wallet-sdk/crypto/base58"
	"github.com/okx/go-wallet-sdk/crypto/bip32"
	bip39 "github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewChildKeyByPathString(t *testing.T) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println(hex.EncodeToString(entropy), mnemonic)

	seed := bip39.NewSeed(mnemonic, "")
	rp, _ := bip32.NewMasterKey(seed)

	c, _ := rp.NewChildKeyByPathString("m/44'/118'/0'/0/0")
	childPrivateKey := hex.EncodeToString(c.Key.Key)
	fmt.Println(childPrivateKey)
}

func TestSignMessage(t *testing.T) {
	f, err := os.OpenFile("./data/bip32.txt", os.O_RDONLY, 0666)
	if err != nil {
		//t.Error(err)
		return
	}

	reader := bufio.NewReader(f)

	ic := 0
	fc := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		ic++
		items := strings.Split(string(line), ",")
		mnemonic := items[0]
		rootPrivateKey := items[1]
		childPrivateKey := items[2]
		secp256k1PublicKey := items[3]
		ed25519PublicKey := items[4]
		msg := items[5]
		secp256k1Signature := items[6]
		ed25519Signature := items[7]

		seed := bip39.NewSeed(mnemonic, "")
		rp, _ := bip32.NewMasterKey(seed)
		rootPrivateKey2 := hex.EncodeToString(rp.Key.Key)
		if rootPrivateKey2 != rootPrivateKey {
			t.Error("rootPrivateKey not match", rootPrivateKey, rootPrivateKey2)
			fc++
			continue
		}
		// m/44'/0'/0'/0/0
		c, _ := rp.NewChildKeyByPath(bip32.FirstHardenedChild+44, 0|bip32.FirstHardenedChild, bip32.FirstHardenedChild, 0, 0)
		childPrivateKey2 := hex.EncodeToString(c.Key.Key)
		if childPrivateKey2 != childPrivateKey {
			t.Error("childPrivateKey not match", childPrivateKey, childPrivateKey2)
			fc++
			continue
		}

		pk, _ := hex.DecodeString(childPrivateKey)
		_, pbk := btcec.PrivKeyFromBytes(pk)
		secp256k1PublicKey2 := hex.EncodeToString(pbk.SerializeCompressed())
		if secp256k1PublicKey2 != secp256k1PublicKey {
			t.Error("secp256k1PublicKey not match", secp256k1PublicKey, secp256k1PublicKey2)
			fc++
			continue
		}

		// The 32-bit private key here is actually the seed
		ed25519PublicKey2 := hex.EncodeToString(ed25519.NewKeyFromSeed(pk).Public().(ed25519.PublicKey))
		if ed25519PublicKey2 != ed25519PublicKey {
			t.Error("ed25519PublicKey not match", ed25519PublicKey, ed25519PublicKey2)
			fc++
			continue
		}

		msgHash, _ := hex.DecodeString(msg)
		edkey := ed25519.NewKeyFromSeed(pk)
		s, _ := edkey.Sign(crypto_rand.Reader, msgHash, crypto.Hash(0))
		ed25519Signature2 := hex.EncodeToString(s)
		if ed25519Signature2 != ed25519Signature {
			t.Error("ed25519Signature not match", ed25519Signature, ed25519Signature2)
			fc++
			continue
		}

		k, _ := btcec.PrivKeyFromBytes(pk)
		s2, _ := ecdsa.SignCompact(k, msgHash, false)
		secp256k1Signature2 := hex.EncodeToString(s2)
		if secp256k1Signature2 != secp256k1Signature {
			t.Error("secp256k1Signature not match", secp256k1Signature, secp256k1Signature2)
			fc++
			continue
		}
	}
	fmt.Printf("The verification succeeds(%d) and fails (%d)\n", ic, fc)
}

func TestHash(t *testing.T) {
	f, err := os.OpenFile("./data/hash.txt", os.O_RDONLY, 0666)
	if err != nil {
		//t.Error(err)
		return
	}

	reader := bufio.NewReader(f)

	ic := 0
	fc := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		ic++
		items := strings.Split(string(line), ",")

		//const msg = base.randomBytes(32)
		//const s1 = base.toHex(base.sha256(msg))
		//const s2 = base.toHex(base.sha512(msg))
		//const s3 = base.toBase58(msg)
		//const s4 = base.toBase58Check(msg)
		//const s5 = base.toBase64(msg)
		//const s6 = base.toBech32("prefix", msg)
		//const s7 = base.toHex(base.hash160(msg))
		//const s8 = base.toHex(base.keccak256(msg))
		//const s9 = base.toHex(base.hmacSHA256("key", msg))
		msg := items[0]
		// sha256
		s1 := items[1]
		// sha512
		s2 := items[2]
		// base58
		s3 := items[3]
		// base58check
		s4 := items[4]
		// base64
		s5 := items[5]
		// bech32
		s6 := items[6]
		// hash160
		s7 := items[7]
		// keccak256
		s8 := items[8]
		// hmacSHA256
		s9 := items[9]

		m, _ := hex.DecodeString(msg)
		t1 := sha256.Sum256(m)
		s11 := hex.EncodeToString(t1[:])
		if s11 != s1 {
			t.Error("sha256 not match", s1, s11)
			fc++
			continue
		}

		t2 := sha512.Sum512(m)
		s22 := hex.EncodeToString(t2[:])
		if s22 != s2 {
			t.Error("sha512 not match", s2, s22)
			fc++
			continue
		}

		s33 := base58.Encode(m)
		if s33 != s3 {
			t.Error("base58 not match", s3, s33)
			fc++
			continue
		}

		s44 := base58.CheckEncodeRaw(m)
		if s44 != s4 {
			t.Error("base58Check not match", s4, s44)
			fc++
			continue
		}

		s55 := base64.StdEncoding.EncodeToString(m)
		if s55 != s5 {
			t.Error("base64 not match", s5, s55)
			fc++
			continue
		}

		s66, _ := bech32.EncodeFromBase256("prefix", m)
		if s66 != s6 {
			t.Error("bech32 not match", s6, s66)
			fc++
			continue
		}

		s77 := hex.EncodeToString(btcutil.Hash160(m))
		if s77 != s7 {
			t.Error("hash160 not match", s7, s77)
			fc++
			continue
		}

		s88 := hex.EncodeToString(keccak256(m))
		if s88 != s8 {
			t.Error("keccak256 not match", s8, s88)
			fc++
			continue
		}

		s99 := hex.EncodeToString(hmacSHA256(m))
		if s99 != s9 {
			t.Error("hmacSHA256 not match", s9, s99)
			fc++
			continue
		}
	}

	fmt.Printf("The verification succeeds(%d) and fails (%d)\n", ic, fc)
}

func keccak256(buf []byte) []byte {
	hash256 := sha3.NewLegacyKeccak256()
	hash256.Write(buf)
	return hash256.Sum(nil)
}

func hmacSHA256(buf []byte) []byte {
	hm := hmac.New(sha256.New, []byte("key"))
	hm.Write(buf)
	return hm.Sum(nil)
}
