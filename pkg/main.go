package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"github.com/sonr-io/core/pkg/ui"
	//	"github.com/sonr-io/core/bind"
)

func main() {
	onExit := func() {
		now := time.Now()
		log.Println(now)
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	// node := sonr.NewNode(reqBytes []byte, call sonr.Callback)
	systray.SetTemplateIcon(ui.GetIcon(ui.SystemTray), ui.GetIcon(ui.SystemTray))
	systray.SetTitle("Sonr")
	systray.SetTooltip("Lantern")

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(ui.GetIcon(ui.SystemTray), ui.GetIcon(ui.SystemTray))
		systray.SetTitle("")
		systray.SetTooltip("Sonr")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItemCheckbox("Unchecked", "Check Me", true)
		mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		// Sets the icon of a menu item. Only available on Mac.
		mEnabled.SetTemplateIcon(ui.GetIcon(ui.SystemTray), ui.GetIcon(ui.SystemTray))

		systray.AddMenuItem("Ignored", "Ignored")
		subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
		subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
		subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
		subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")

		mUrl := systray.AddMenuItem("Open UI", "my home")
		systray.AddSeparator()
		mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
		go func() {
			<-mQuitOrig.ClickedCh
			fmt.Println("Requesting quit")
			systray.Quit()
			fmt.Println("Finished quitting")
		}()
		shown := true
		toggle := func() {
			if shown {
				subMenuBottom.Check()
				subMenuBottom2.Hide()
				mQuitOrig.Hide()
				mEnabled.Hide()
				shown = false
			} else {
				subMenuBottom.Uncheck()
				subMenuBottom2.Show()
				mQuitOrig.Show()
				mEnabled.Show()
				shown = true
			}
		}

		for {
			select {
			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mEnabled.ClickedCh:
				mEnabled.SetTitle("Disabled")
				mEnabled.Disable()
			case <-mUrl.ClickedCh:
				err := open.Run("https://www.getlantern.org")
				if err != nil {
					log.Fatalln(err)
				}
			case <-subMenuBottom2.ClickedCh:
				panic("panic button pressed")
			case <-subMenuBottom.ClickedCh:
				toggle()
			}
		}
	}()
}
