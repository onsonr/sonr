package ctx

// ╭───────────────────────────────────────────────────────────╮
// │                      Request Headers                      │
// ╰───────────────────────────────────────────────────────────╯

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

// ╭───────────────────────────────────────────────────────────╮
// │                     Response Headers                      │
// ╰───────────────────────────────────────────────────────────╯

type ResponseHeaders struct {
	// HTMX Specific
	HXLocation           *string `header:"HX-Location"`
	HXPushURL            *string `header:"HX-Push-Url"`
	HXRedirect           *string `header:"HX-Redirect"`
	HXRefresh            *string `header:"HX-Refresh"`
	HXReplaceURL         *string `header:"HX-Replace-Url"`
	HXReswap             *string `header:"HX-Reswap"`
	HXRetarget           *string `header:"HX-Retarget"`
	HXReselect           *string `header:"HX-Reselect"`
	HXTrigger            *string `header:"HX-Trigger"`
	HXTriggerAfterSettle *string `header:"HX-Trigger-After-Settle"`
	HXTriggerAfterSwap   *string `header:"HX-Trigger-After-Swap"`
}

type ProtectedResponseHeaders struct {
	AcceptCH                      *string `header:"Accept-CH"`
	AccessControlAllowCredentials *string `header:"Access-Control-Allow-Credentials"`
	AccessControlAllowHeaders     *string `header:"Access-Control-Allow-Headers"`
	AccessControlAllowMethods     *string `header:"Access-Control-Allow-Methods"`
	AccessControlExposeHeaders    *string `header:"Access-Control-Expose-Headers"`
	AccessControlRequestHeaders   *string `header:"Access-Control-Request-Headers"`
	ContentSecurityPolicy         *string `header:"Content-Security-Policy"`
	CrossOriginEmbedderPolicy     *string `header:"Cross-Origin-Embedder-Policy"`
	PermissionsPolicy             *string `header:"Permissions-Policy"`
	ProxyAuthorization            *string `header:"Proxy-Authorization"`
	WWWAuthenticate               *string `header:"WWW-Authenticate"`
}
