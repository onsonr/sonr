package ui

import (
	"log"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr"
	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	win "github.com/sonr-io/core/pkg/window"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type AppInterface struct {
	node       *sonr.Node
	box        packr.Box
	mLink      *systray.MenuItem
	mCount     *systray.MenuItem
	mQuit      *systray.MenuItem
	mPeersList []*systray.MenuItem
	peerCount  int32
	lobbySize  int32
}

// ^ Start Starts System tray with Library ^ //
func Start(box packr.Box) AppInterface {
	// Check Platform
	if runtime.GOOS == "windows" {
		icon, err := box.Find("systray.ico")
		if err != nil {
			log.Println(err)
		}
		systray.SetIcon(icon)
	} else {
		icon, err := box.Find("systray.png")
		if err != nil {
			log.Println(err)
		}
		systray.SetTemplateIcon(icon, icon)
	}

	// Set Initial Menu Vars
	systray.SetTooltip("Sonr")

	// Default
	ai := AppInterface{
		peerCount: 0,
		lobbySize: 1,
		box:       box,
	}

	// Link Sonr Device
	ai.mLink = systray.AddMenuItem("Link Device", "Link a Device to Sonr")
	ai.mLink.SetTemplateIcon(ai.GetIcon(Link), ai.GetIcon(Link))

	// Quit Sonr
	ai.mQuit = systray.AddMenuItem("Quit", "Quit Sonr Desktop")
	ai.mQuit.SetTemplateIcon(ai.GetIcon(Close), ai.GetIcon(Close))
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
				// Create Link Request
				name := ShowNameDialog()
				linkRequest := ai.node.LinkRequest(name)

				// Get Node JSON
				jsonBytes, err := protojson.Marshal(linkRequest)
				if err != nil {
					log.Panicln(err)
				}

				// Display Window
				go win.OpenQRWindow(string(jsonBytes))
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

				// Add Files
				files := make([]*md.InviteRequest_FileInfo, 0)
				files = append(files, &md.InviteRequest_FileInfo{
					Path: filename,
				})

				// Create Request
				procReq := md.InviteRequest{
					To:    peer,
					Files: files,
					Type:  md.InviteRequest_File,
				}

				// Convert to Bytes
				reqBytes, err := proto.Marshal(&procReq)
				if err != nil {
					log.Panicln(err)
				}

				ai.node.Invite(reqBytes)
			} else {
				log.Println("Node not set.")
			}

			// @ Link Item Clicked
		case <-linkItem.ClickedCh:
			// Validate and Invite URL
			if ai.node != nil {
				url := ShowURLDialog()

				// Create Request
				procReq := md.InviteRequest{
					To:   peer,
					Url:  url,
					Type: md.InviteRequest_File,
				}

				// Convert to Bytes
				reqBytes, err := proto.Marshal(&procReq)
				if err != nil {
					log.Panicln(err)
				}

				ai.node.Invite(reqBytes)
			} else {
				log.Println("Node not set.")
			}
		}
	}
}

// ^ Method to Rebuild Menu for Lobby Refresh ^ //
func (ai *AppInterface) RefreshPeers(newLob *md.Lobby, node *sonr.Node) {
	// Set Node
	ai.node = node

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
	peerItem := systray.AddMenuItem(p.Profile.FirstName, "")
	peerItem.SetTemplateIcon(ai.GetDeviceIcon(p.Platform), ai.GetDeviceIcon(p.Platform))

	// Add Peer Send Options
	urlItem := peerItem.AddSubMenuItem("Send URL", "Send a URL to "+p.Profile.FirstName)
	urlItem.SetTemplateIcon(ai.GetIcon(URL), ai.GetIcon(URL))
	fileItem := peerItem.AddSubMenuItem("Send File", "Send a File to "+p.Profile.FirstName)
	fileItem.SetTemplateIcon(ai.GetIcon(File), ai.GetIcon(File))

	// Spawn Routine to handle Peer Item Actions
	go ai.HandlePeerInput(fileItem, urlItem, p)

	// Add to Menu List
	ai.mPeersList = append(ai.mPeersList, peerItem)
}
