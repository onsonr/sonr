package ctx

import "github.com/go-webauthn/webauthn/protocol"

type WebBytes = protocol.URLEncodedBase64

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
