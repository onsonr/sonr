package textile

import (
	"context"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Method Reads Inbox and Returns List of Mail Entries
func (tn *textile) ReadMail() ([]*md.MailEntry, *md.SonrError) {
	// Check Mail Enabled
	if tn.active && tn.options.GetMailbox() {
		// List the recipient's inbox
		inbox, err := tn.mailbox.ListInboxMessages(context.Background())

		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Initialize Entry List
		entries := make([]*md.MailEntry, len(inbox))

		// Iterate over Entries
		for i, v := range inbox {
			// Open decrypts the message body
			body, err := v.Open(context.Background(), tn.identity)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Unmarshal Body to entry
			entry := &md.MailEntry{}
			err = proto.Unmarshal(body, entry)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}
			entries[i] = entry
		}
		return entries, nil
	}
	return nil, nil
}

// @ Method Sends Mail Entry to Peer
func (tn *textile) SendMail(e *md.MailEntry) *md.SonrError {
	// Check Mail Enabled
	if tn.active && tn.options.GetMailbox() {
		// Send Message to Mailbox
		_, err := tn.mailbox.SendMessage(context.Background(), e.ToPubKey(), e.Buffer())

		// Check Error
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}
	}
	return nil
}
