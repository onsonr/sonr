// TODO: Update this Package to utlize: https://github.com/go-webauthn/example

package webauthn

import (
	"fmt"

	wan "github.com/go-webauthn/webauthn/webauthn"
)

func Init() error {
	wan, err := wan.New(&wan.Config{
		RPDisplayName: "Go Webauthn",                        // Display Name for your site
		RPID:          "go-webauthn.local",                  // Generally the FQDN for your site
		RPOrigin:      "https://login.go-webauthn.local",    // The origin URL for WebAuthn requests
		RPIcon:        "https://go-webauthn.local/logo.png", // Optional icon URL for your site
	})
	if err != nil {
		fmt.Println(err)
	}
	webAuthn = wan
	return nil
}
