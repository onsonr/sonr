// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

import "github.com/apple/pkl-go/pkl"

// Represents a transaction body
type TxBody struct {
	Messages []Msg `pkl:"messages"`

	Memo *string `pkl:"memo"`

	TimeoutHeight *int `pkl:"timeoutHeight"`

	ExtensionOptions *[]*pkl.Object `pkl:"extensionOptions"`

	NonCriticalExtensionOptions *[]*pkl.Object `pkl:"nonCriticalExtensionOptions"`
}
