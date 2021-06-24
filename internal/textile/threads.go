package textile

import (
	"fmt"
	"log"

	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/core/thread"
)

func (tn *textile) InitThreads() *md.SonrError {
	// Check Thread Enabled
	if tn.active && tn.options.GetThreads() {
		// Generate a new thread ID
		threadID := thread.NewIDV1(thread.Raw, 32)

		// Create your new thread
		err := tn.client.NewDB(tn.ctxToken, threadID)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Get DB Info
		info, err := tn.client.GetDBInfo(tn.ctxToken, threadID)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Log DB Info
		log.Println("> Success!")
		log.Println(fmt.Sprintf("ID: %s \n Maddr: %s \n Key: %s \n Name: %s \n", threadID.String(), info.Addrs, info.Key.String(), info.Name))
	}
	return nil
}
