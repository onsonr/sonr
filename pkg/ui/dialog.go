package ui

import (
	"log"

	"github.com/gen2brain/dlgs"
	md "github.com/sonr-io/core/internal/models"
)

const kCompatibleFileTypes = "*.png *.jpg *.jpeg *.mp4 *.avi"

// ^ Presents a Authentication Dialog for Approval ^ //
func ShowAuthDialog(inv *md.AuthInvite) bool {
	// Set Text
	description := "Do you accept " + inv.From.Profile.FirstName + "'s invite to receive an " + inv.Payload.String()

	// Display
	decision, err := dlgs.Question("Sonr Invitation", description, true)
	if err != nil {
		log.Panicln(err)
	}

	// Return
	return decision
}

// ^ Presents a Error Message ^ //
func ShowErrorDialog(errorMsg *md.ErrorMessage) {
	_, err := dlgs.Error("Error", errorMsg.Message)
	if err != nil {
		panic(err)
	}
}

// ^ Presents a File Input Dialog ^ //
func ShowFileDialog() string {
	// Display
	filename, _, err := dlgs.File("Select a File to Send...", kCompatibleFileTypes, false)
	if err != nil {
		log.Panicln(err)
	}

	// Return
	return filename
}

// ^ Presents a URL Input Dialog ^ //
func ShowURLDialog() string {
	// Display
	url, _, err := dlgs.Entry("URL Link", "Enter a URL Here: ", "")
	if err != nil {
		log.Println(err)
	}

	// Return
	return url
}
