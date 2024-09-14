// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

// Represents a transaction body
type TxBody struct {
	Messages []Msg `pkl:"messages"`

	Memo *string `pkl:"memo"`

	TimeoutHeight *int `pkl:"timeoutHeight"`

	ExtensionOptions *[]*pkl.Object `pkl:"extensionOptions"`

	NonCriticalExtensionOptions *[]*pkl.Object `pkl:"nonCriticalExtensionOptions"`
}
