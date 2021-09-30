package mailbox

type MailboxProtocol struct {
	
}
// // Method Sends Mail Entry to Peer
// func (sc *serviceClient) SendMail(inv *data.InviteRequest) *data.SonrError {
// 	// Check Mail Enabled
// 	if sc.Textile.options.GetMailbox() {
// 		// Fetch Peer Thread Key
// 		pubKey, serr := inv.GetTo().GetActive().ThreadKey()
// 		if serr != nil {
// 			return serr
// 		}

// 		// Marshal Data
// 		buf, err := protojson.Marshal(inv)
// 		if err != nil {
// 			return data.NewMarshalError(err)
// 		}

// 		// Send to Mailbox
// 		serr = sc.Textile.sendMail(pubKey, buf)
// 		if serr != nil {
// 			return serr
// 		}
// 		logger.Info("Succesfully sent mail!")
// 		return nil
// 	}
// 	return nil
// }

// // Method Handles a given Mailbox Request for a Message
// func (sc *serviceClient) HandleMailbox(req *data.MailboxRequest) (*data.MailboxResponse, *data.SonrError) {
// 	if req.Action == data.MailboxRequest_READ {
// 		// Set Mailbox Message as Read
// 		err := sc.Textile.readMessage(req.ID)
// 		if err != nil {
// 			return &data.MailboxResponse{
// 				Success: false,
// 				Action:  data.MailboxResponse_Action(req.Action),
// 			}, err
// 		}

// 		// Return Success
// 		return &data.MailboxResponse{
// 			Success: true,
// 			Action:  data.MailboxResponse_Action(req.Action),
// 		}, nil
// 	} else if req.Action == data.MailboxRequest_DELETE {
// 		// Delete Mailbox Message
// 		err := sc.Textile.deleteMessage(req.ID)
// 		if err != nil {
// 			return &data.MailboxResponse{
// 				Success: false,
// 				Action:  data.MailboxResponse_Action(req.Action),
// 			}, err
// 		}
// 		return &data.MailboxResponse{
// 			Success: true,
// 			Action:  data.MailboxResponse_Action(req.Action),
// 		}, nil
// 	} else {
// 		return &data.MailboxResponse{
// 			Success: false,
// 			Action:  data.MailboxResponse_Action(req.Action),
// 		}, data.NewErrorWithType(data.ErrorEvent_MAILBOX_ACTION_INVALID)
// 	}
// }
