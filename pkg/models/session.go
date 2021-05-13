package models

func NewOutSession(p *Peer, req *InviteRequest, tc NodeCallback) *Session {
	return &Session{
		File:      req.GetFile(),
		Receiver:  req.GetTo(),
		Sender:    p,
		Index:     0,
		Direction: Session_Outgoing,
	}
}
