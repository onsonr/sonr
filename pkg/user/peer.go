package user

import (
	"hash/fnv"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/internal/models"
)



// ^ GetPeerID returns ID Reference ^ //
func GetPeerID(device *md.Device, profile *md.Profile, peerID string) *md.Peer_ID {
	// Initialize
	deviceID := device.GetId()

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.GetUsername()))
	if err != nil {
		return nil
	}

	// Check if ID not provided
	if deviceID == "" {
		// Generate ID
		if id, err := mid.ProtectedID("Sonr"); err != nil {
			sentry.CaptureException(err)
			deviceID = ""
		} else {
			deviceID = id
		}
	}

	return &md.Peer_ID{
		Peer:   peerID,
		Device: deviceID,
		User:   userID.Sum32(),
	}
}
