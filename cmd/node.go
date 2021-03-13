package sonr

import (
	"log"
	"math"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Update proximity/direction and Notify Lobby ^ //
func (sn *Node) Update(facing float64, heading float64) {
	// Update User Values
	var faceDir float64
	var faceAnpd float64
	var headDir float64
	var headAnpd float64
	faceDir = math.Round(facing*100) / 100
	headDir = math.Round(heading*100) / 100
	desg := int((facing / 11.25) + 0.25)

	// Find Antipodal
	if facing > 180 {
		faceAnpd = math.Round((facing-180)*100) / 100
	} else {
		faceAnpd = math.Round((facing+180)*100) / 100
	}

	// Find Antipodal
	if heading > 180 {
		headAnpd = math.Round((heading-180)*100) / 100
	} else {
		headAnpd = math.Round((heading+180)*100) / 100
	}

	// Set Position
	sn.peer.Position = &md.Position{
		Facing:           faceDir,
		FacingAntipodal:  faceAnpd,
		Heading:          headDir,
		HeadingAntipodal: headAnpd,
		Designation:      md.Position_Designation(desg % 32),
	}

	// Inform Lobby
	err := sn.lobby.Update()
	if err != nil {
		sn.error(err, "Update")
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (sn *Node) Invite(reqBytes []byte) {
	// @ 1. Initialize from Request
	req := &md.InviteRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		log.Println(err)
	}

	// @ 2. Check Transfer Type
	if req.Type == md.InviteRequest_Contact || req.Type == md.InviteRequest_URL {
		// @ 3. Send Invite to Peer
		// Set Contact
		req.Contact = sn.contact
		invMsg := md.NewInviteFromRequest(req, sn.peer)

		// Get PeerID and Check error
		id, _, err := sn.lobby.Find(req.To.Id.Peer)
		if err != nil {
			sn.error(err, "InviteWithContact")
		}

		// Run Routine
		go func(inv *md.AuthInvite) {
			// Convert Protobuf to bytes
			msgBytes, err := proto.Marshal(inv)
			if err != nil {
				sn.error(err, "Marshal")
			}

			sn.peerConn.Request(sn.host, id, msgBytes)
		}(&invMsg)
	} else {
		// File Transfer
		sn.queue.AddFromRequest(req)
	}

	// Update Status
	sn.status = md.Status_PENDING
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// Send Response on PeerConnection
	sn.peerConn.Authorize(decision, sn.contact, sn.peer)

	// Update Status
	if decision {
		sn.status = md.Status_INPROGRESS
	} else {
		sn.status = md.Status_AVAILABLE
	}
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Pause() {
	// Check if Response Is Invited
	if sn.status == md.Status_INVITED {
		sn.peerConn.Cancel(sn.peer)
	}
	err := sn.lobby.Standby()
	if err != nil {
		sn.error(err, "Pause")
	}
	md.GetState().Pause()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Resume() {
	err := sn.lobby.Resume()
	if err != nil {
		sn.error(err, "Resume")
	}
	md.GetState().Resume()
}

// ^ Close Ends All Network Communication ^
func (sn *Node) Stop() {
	// Check if Response Is Invited
	if sn.status == md.Status_INVITED {
		sn.peerConn.Cancel(sn.peer)
	}
	sn.host.Close()
}
