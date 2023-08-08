package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/types/tx"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  SignAndSendTx                                 ||
// ! ||--------------------------------------------------------------------------------||

// SignAndSendTxRequest struct is used to represent a request object for signing and sending a transaction. the messages to be included in the transaction, and `DID`, which is a string representing the decentralized identifier associated with the transaction.
type SignAndSendTxRequest struct {
	Messages []*tx.SignDoc `json:"messages"`
	DID      string        `json:"did"`
}

// Marshal returns the JSON encoding of the SignAndSendTxRequest object.
func (r *SignAndSendTxRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal parses the JSON-encoded data and stores the result in the SignAndSendTxRequest object.
func (r *SignAndSendTxRequest) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 SignWithAccount                                ||
// ! ||--------------------------------------------------------------------------------||

// SignWithAccountRequest struct is defining the structure of a request object for signing a message with an account. It has two fields: `Message`, which is a string representing the message to be signed, and `DID`, which is a string representing the decentralized identifier
// associated with the account.
type SignWithAccountRequest struct {
	Message string `json:"message"`
	DID     string `json:"did"`
}

// Marshal returns the JSON encoding of the SignWithAccountRequest object.
func (r *SignWithAccountRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal parses the JSON-encoded data and stores the result in the SignWithAccountRequest object.
func (r *SignWithAccountRequest) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                VerifyWithAccount                               ||
// ! ||--------------------------------------------------------------------------------||

// VerifyWithAccountRequest struct is defining the structure of a request object for verifying a message with an account. It has three fields: message, and `DID`, which is a string representing the decentralized identifier associated with the account.
type VerifyWithAccountRequest struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	DID       string `json:"did"`
}

// Marshal returns the JSON encoding of the VerifyWithAccountRequest object.
func (r *VerifyWithAccountRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal parses the JSON-encoded data and stores the result in the VerifyWithAccountRequest object.
func (r *VerifyWithAccountRequest) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
