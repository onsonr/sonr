package headers

type RequestHeaders struct {
	CacheControl  *string `header:"Cache-Control"`
	DeviceMemory  *string `header:"Device-Memory"`
	From          *string `header:"From"`
	Host          *string `header:"Host"`
	Referer       *string `header:"Referer"`
	UserAgent     *string `header:"User-Agent"`
	ViewportWidth *string `header:"Viewport-Width"`
	Width         *string `header:"Width"`

	// HTMX Specific
	HXBoosted               *string `header:"HX-Boosted"`
	HXCurrentURL            *string `header:"HX-Current-URL"`
	HXHistoryRestoreRequest *string `header:"HX-History-Restore-Request"`
	HXPrompt                *string `header:"HX-Prompt"`
	HXRequest               *string `header:"HX-Request"`
	HXTarget                *string `header:"HX-Target"`
	HXTriggerName           *string `header:"HX-Trigger-Name"`
	HXTrigger               *string `header:"HX-Trigger"`
}

type ProtectedRequestHeaders struct {
	Authorization      *string `header:"Authorization"`
	Forwarded          *string `header:"Forwarded"`
	Link               *string `header:"Link"`
	PermissionsPolicy  *string `header:"Permissions-Policy"`
	ProxyAuthorization *string `header:"Proxy-Authorization"`
	WWWAuthenticate    *string `header:"WWW-Authenticate"`
}
