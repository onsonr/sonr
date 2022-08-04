---
title: Create Cosmos Wallet in Go
id: setup-wallet
displayed_sidebar: runSidebar
---

## New Cosmos Wallet

### Wallet Options
```go

// WalletOption is a function that modifies the options for a Wallet.
type WalletOption func(*options) error

// WithPassphrase sets the passphrase for the keyring.
func WithPassphrase(s string) WalletOption {
	return func(o *options) error {
		o.passphrase = s
		return nil
	}
}

// WithFolderPath sets the folder path for the keyring.
func WithFolderPath(s string) WalletOption {
	return func(o *options) error {
		o.folder = device.Folder(s)
		return nil
	}
}

// WithWalletName sets the name of the wallet to be created.
func WithWalletName(s string) WalletOption {
	return func(o *options) error {
		o.walletName = s
		return nil
	}
}

type options struct {
	walletName string
	passphrase string
	folder     device.Folder
}

```


### Create Wallet
```go

g := defaultOptions()
	for _, option := range options {
		if err := option(g); err != nil {
			return nil, "", err
		}
	}

	// Add keys and see they return in alphabetical order
	_, mnemonic, err := kr.NewMnemonic(g.walletName, keyring.English, sdk.FullFundraiserPath, g.passphrase, hd.Secp256k1)
	if err != nil {
		return nil, "", err
	}

	// Create default sonr key
	ks, err := CreateKeySet(mnemonic)
	if err != nil {
		return nil, "", err
	}

	// Copy keys to keyring if not already there
	_, err = ks.CopyToKeyring(kr, g.walletName)
	if err == nil {
		err = ks.Export(device.Support)
		if err != nil {
			return nil, "", err
		}
	}

	return ks, mnemonic, nil
```

## Export Cosmos Wallet

```go

// ExportWallet returns armored private key and public key
func ExportWallet(kr keyring.Keyring, name string, passphrase string) (string, error) {
	armor, err := kr.ExportPrivKeyArmor(name, passphrase)
	if err != nil {
		return "", err
	}
	return armor, nil
}
```

## Restore Cosmos Wallet

```go

// RestoreWallet restores a private key from ASCII armored format.
func RestoreWallet(name string, armor string, passphrase string) (keyring.Keyring, error) {
	privKey, algo, err := crypto.UnarmorDecryptPrivKey(armor, passphrase)
	if err != nil {
		return nil, err
	}
	kr := keyring.NewInMemory()
	if err := kr.ImportPrivKey(name, algo, passphrase); err != nil {
		return nil, err
	}
	log.Println(privKey.PubKey())
	return kr, nil
}
```
