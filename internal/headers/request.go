package headers

type RequestHeaders struct {
	Authorization      *string `header:"Authorization"`
	CacheControl       *string `header:"Cache-Control"`
	DeviceMemory       *string `header:"Device-Memory"`
	Forwarded          *string `header:"Forwarded"`
	From               *string `header:"From"`
	Host               *string `header:"Host"`
	Link               *string `header:"Link"`
	PermissionsPolicy  *string `header:"Permissions-Policy"`
	ProxyAuthorization *string `header:"Proxy-Authorization"`
	Referer            *string `header:"Referer"`
	UserAgent          *string `header:"User-Agent"`
	ViewportWidth      *string `header:"Viewport-Width"`
	Width              *string `header:"Width"`
	WWWAuthenticate    *string `header:"WWW-Authenticate"`

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
