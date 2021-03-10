package window

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/progrium/macdriver/webkit"
)

func NewMacWindow() {
	app := cocoa.NSApp_WithDidLaunch(func(notification objc.Object) {
		config := webkit.WKWebViewConfiguration_New()
		wv := webkit.WKWebView_Init(core.Rect(0, 0, 1440, 900), config)
		url := core.URL("http://progrium.com")
		req := core.NSURLRequest_Init(url)
		wv.LoadRequest(req)

		w := cocoa.NSWindow_Init(core.Rect(0, 0, 1440, 900),
			cocoa.NSClosableWindowMask|
				cocoa.NSTitledWindowMask,
			cocoa.NSBackingStoreBuffered, false)
		w.SetContentView(wv)
		w.MakeKeyAndOrderFront(w)
		w.Center()
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
