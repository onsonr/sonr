package models

import "fmt"

func (r *RemoteCreateRequest) GetTopic() string {
	return fmt.Sprintf("%s.remote.%s.snr/", r.Fingerprint, r.SName)
}

// Get Remote Point Info
func (r *RemoteCreateRequest) NewCreatedRemote(u *User) *Lobby {
	// Create Lobby
	return &Lobby{
		// General
		Type:  Lobby_REMOTE,
		Peers: make(map[string]*Peer),
		User:  u.GetPeer(),

		// Info
		Info: &Lobby_Remote{
			Remote: &Lobby_RemoteInfo{
				IsJoin:      false,
				Fingerprint: r.GetFingerprint(),
				Words:       r.GetWords(),
				Topic:       r.GetTopic(),
				File:        r.GetFile(),
				Owner:       u.GetPeer(),
			},
		},
	}
}

func (r *RemoteJoinRequest) NewJoinedRemote(u *User) *Lobby {
	// Create Lobby
	return &Lobby{
		// General
		Type:  Lobby_REMOTE,
		Peers: make(map[string]*Peer),
		User:  u.GetPeer(),

		// Info
		Info: &Lobby_Remote{
			Remote: &Lobby_RemoteInfo{
				IsJoin: true,
				Topic:  r.GetTopic(),
			},
		},
	}
}
