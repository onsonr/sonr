// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

import "github.com/apple/pkl-go/pkl"

func init() {
	pkl.RegisterMapping("vault", Vault{})
	pkl.RegisterMapping("vault#Sonr", SonrImpl{})
	pkl.RegisterMapping("vault#Sqlite", SqliteImpl{})
	pkl.RegisterMapping("vault#Account", AccountImpl{})
	pkl.RegisterMapping("vault#Asset", AssetImpl{})
	pkl.RegisterMapping("vault#Chain", ChainImpl{})
	pkl.RegisterMapping("vault#Credential", CredentialImpl{})
	pkl.RegisterMapping("vault#Profile", ProfileImpl{})
	pkl.RegisterMapping("vault#Property", PropertyImpl{})
	pkl.RegisterMapping("vault#Keyshare", KeyshareImpl{})
	pkl.RegisterMapping("vault#PublicKey", PublicKeyImpl{})
	pkl.RegisterMapping("vault#Permission", PermissionImpl{})
}
