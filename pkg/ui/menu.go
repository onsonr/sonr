package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/getlantern/systray"
	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type AppInterface struct {
	node       *sonr.Node
	mLink      *systray.MenuItem
	mCount     *systray.MenuItem
	mQuit      *systray.MenuItem
	mPeersList []*systray.MenuItem
	peerCount  int32
	lobbySize  int32
	App        fyne.App
}

// ^ Start Starts System tray with Library ^ //
func Start() AppInterface {
	// Set Initial Menu Vars
	systray.SetTemplateIcon(GetIcon(SystemTray), GetIcon(SystemTray))
	systray.SetTooltip("Sonr")

	// Default
	ai := AppInterface{
		peerCount: 0,
		lobbySize: 1,
		App:       app.New(),
	}
	ai.peerCount = 0
	ai.lobbySize = 1

	// Link Sonr Device
	ai.mLink = systray.AddMenuItem("Link Device", "Link a Device to Sonr")
	ai.mLink.SetTemplateIcon(GetIcon(Link), GetIcon(Link))

	// Quit Sonr
	ai.mQuit = systray.AddMenuItem("Quit", "Quit Sonr Desktop")
	ai.mQuit.SetTemplateIcon(GetIcon(Close), GetIcon(Close))
	systray.AddSeparator()

	// Pers Label
	ai.mCount = systray.AddMenuItem("Available Peers", "")
	ai.mCount.Disable()
	return ai
}

// ^ References Sonr Node and Handles Input ^ //
func (ai *AppInterface) Initialize(n *sonr.Node) {
	// Set Node
	ai.node = n

	// Handle Menu Events
	go ai.HandleMenuInput()
}

// ^ Routine Handles Menu Input ^ //
func (ai *AppInterface) HandleMenuInput() {
	for {
		select {
		// @ File Item Clicked
		case <-ai.mQuit.ClickedCh:
			systray.Quit()

			// @ Link Item Clicked
		case <-ai.mLink.ClickedCh:
			// Validate Node Set
			if ai.node != nil {
				// Get Node JSON
				jsonBytes, err := protojson.Marshal(ai.node.Peer())
				if err != nil {
					log.Panicln(err)
				}

				// Display Window
				ai.OpenQRWindow(string(jsonBytes))
			} else {
				log.Println("Node not set.")
			}
		}
	}
}

// ^ Routine Handles Peer Item Input ^ //
func (ai *AppInterface) HandlePeerInput(fileItem *systray.MenuItem, linkItem *systray.MenuItem, peer *md.Peer) {
	for {
		select {
		// @ File Item Clicked
		case <-fileItem.ClickedCh:
			// Validate and Invite File
			if ai.node != nil {
				// Get File
				filename := ShowFileDialog()

				// Create Request
				procReq := md.ProcessRequest{
					IsExternal: false,
					FilePath:   filename,
				}

				// Convert to Bytes
				reqBytes, err := proto.Marshal(&procReq)
				if err != nil {
					log.Panicln(err)
				}

				ai.node.Process(reqBytes)
				ai.node.InviteWithFile(peer.Id)
			} else {
				log.Println("Node not set.")
			}

			// @ Link Item Clicked
		case <-linkItem.ClickedCh:
			// Validate and Invite URL
			if ai.node != nil {
				url := ShowURLDialog()
				ai.node.InviteWithURL(peer.Id, url)
			} else {
				log.Println("Node not set.")
			}
		}
	}
}

// ^ Method to Rebuild Menu for Lobby Refresh ^ //
func (ai *AppInterface) RefreshPeers(newLob *md.Lobby) {
	// Check if Lobby Updated
	if newLob.Size != ai.lobbySize {
		// Change Lobby
		ai.lobbySize = newLob.Size
		ai.peerCount = newLob.Size - 1

		// Reset Menu
		for _, mi := range ai.mPeersList {
			mi.Hide()
		}

		// Empty Peers
		ai.mPeersList = nil

		// Add Peers to Menu
		for _, p := range newLob.Peers {
			ai.SetPeerItem(p)
		}
	}
}

// ^ Method to Build Peer Item ^ //
func (ai *AppInterface) SetPeerItem(p *md.Peer) {
	// Add Peer to Menu
	peerItem := systray.AddMenuItem(p.FirstName, "")
	peerItem.SetTemplateIcon(GetDeviceIcon(p.Device), GetDeviceIcon(p.Device))

	// Add Peer Send Options
	urlItem := peerItem.AddSubMenuItem("Send URL", "Send a URL to "+p.FirstName)
	urlItem.SetTemplateIcon(GetIcon(URL), GetIcon(URL))
	fileItem := peerItem.AddSubMenuItem("Send File", "Send a File to "+p.FirstName)
	fileItem.SetTemplateIcon(GetIcon(File), GetIcon(File))

	// Spawn Routine to handle Peer Item Actions
	go ai.HandlePeerInput(fileItem, urlItem, p)

	// Add to Menu List
	ai.mPeersList = append(ai.mPeersList, peerItem)
}
