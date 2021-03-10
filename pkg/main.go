package main

import (
	"flag"
	"log"
	"time"

	"context"

	"github.com/getlantern/systray"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	ui "github.com/sonr-io/core/pkg/interface"
	sv "github.com/sonr-io/core/pkg/service"

	md "github.com/sonr-io/core/internal/models"
)

type SysInfo struct {
	OLC       string
	Device    md.Device
	Directory md.Directories
}

// Define Context
var desk *sv.Client
var ctx context.Context
var app ui.Interface

func main() {

	ctx = context.Background()
	systray.Run(onReady, onExit)

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		fontName := flag.String("font", "Helvetica", "font to use")
		screen := cocoa.NSScreen_Main().Frame().Size
		tr, fontSize := func() (rect core.NSRect, size float64) {
			t := cocoa.NSTextView_Init(core.Rect(0, 0, 0, 0))
			t.SetString("Sonr.io")
			for s := 70.0; s <= 550; s += 12 {
				t.SetFont(cocoa.Font(*fontName, s))
				t.LayoutManager().EnsureLayoutForTextContainer(t.TextContainer())
				rect = t.LayoutManager().UsedRectForTextContainer(t.TextContainer())
				size = s
				if rect.Size.Width >= screen.Width*0.8 {
					break
				}
			}
			return rect, size
		}()

		height := tr.Size.Height * 1.5
		tr.Origin.Y = (height / 2) - (tr.Size.Height / 2)
		t := cocoa.NSTextView_Init(tr)
		t.SetString("Sonr.io")
		t.SetFont(cocoa.Font(*fontName, fontSize))
		t.SetEditable(false)
		t.SetImportsGraphics(false)
		t.SetDrawsBackground(false)

		c := cocoa.NSView_Init(core.Rect(0, 0, 0, 0))
		c.SetBackgroundColor(cocoa.Color(0, 0, 0, 0.75))
		c.SetWantsLayer(true)
		c.Layer().SetCornerRadius(32.0)
		c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

		tr.Size.Height = height
		tr.Origin.X = (screen.Width / 2) - (tr.Size.Width / 2)
		tr.Origin.Y = (screen.Height / 2) - (tr.Size.Height / 2)

		w := cocoa.NSWindow_Init(core.Rect(0, 0, 0, 0),
			cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
		w.SetContentView(c)
		w.SetTitlebarAppearsTransparent(true)
		w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		w.SetOpaque(false)
		w.SetBackgroundColor(cocoa.NSColor_Clear())
		w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
		w.SetFrameDisplay(tr, true)
		w.MakeKeyAndOrderFront(nil)

		events := make(chan cocoa.NSEvent, 64)
		go func() {
			<-events
			cocoa.NSApp().Terminate()
		}()
		cocoa.NSEvent_GlobalMonitorMatchingMask(cocoa.NSEventMaskAny, events)
	})
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func onReady() {
	// Starts Menu Bar
	app = ui.Start()

	// Creates New Client
	desk = sv.NewClient(ctx, app)
}

func onExit() {
	log.Println(time.Now())
	ctx.Done()
}
