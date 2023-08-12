package routes

const (
	// Accounts endpoints
	getHealthStatusEndpoint    = "/health"                                        // Public
	createAccountEndpoint      = "/sonr/highway/accounts/create/:coin_type/:name" // Authenticated
	kGetAccountEndpoint        = "/sonr/highway/accounts/:did"                    // Public
	kListAccountsEndpoint      = "/sonr/highway/accounts/:coin_type"              // Authenticated
	kSignWithAccountEndpoint   = "/sonr/highway/accounts/:did/sign"               // Authenticated
	kVerifyWithAccountEndpoint = "/sonr/highway/accounts/:did/verify"             // Public

	// Authentication endpoints
	kCurrentControllerEndpoint       = "/sonr/highway/auth/current"                        // Authenticated
	kGetCredCreationOptionsEndpoint  = "/sonr/highway/auth/:origin/register/start/:alias"  // Public
	kGetCredAssertionOptionsEndpoint = "/sonr/highway/auth/:origin/login/start/:alias"     // Public
	kRegisterCredForClaimsEndpoint   = "/sonr/highway/auth/:origin/register/finish/:alias" // Public
	kVerifyCredForAccountEndpoint    = "/sonr/highway/auth/:origin/login/finish/:alias"    // Public
	kGetMagicEmailStartEndpoint      = "/sonr/highway/auth/:origin/magic"                  // Public
	kRegisterMagicEmailEndpoint      = "/sonr/highway/auth/:origin/magic/validate"         // Public

	// Maibox endpoints
	kGetMailboxMessagesEndpoint    = "/sonr/highway/mailbox/:did"                          // Authenticated
	kOpenMailboxMessageEndpoint    = "/sonr/highway/mailbox/:did/:messageId/read"          // Authenticated
	kSendMailboxMessageEndpoint    = "/sonr/highway/mailbox/:from/:to/send"                // Authenticated
	kDeleteMailboxMessageEndpoint  = "/sonr/highway/mailbox/:did/:messageId/delete"        // Authenticated
	kPublishChannelMessageEndpoint = "/sonr/highway/mailbox/channels/:channelId/publish"   // Authenticated
	kSubscribeChannelEndpoint      = "/sonr/highway/mailbox/channels/:channelId/subscribe" // Authenticated
)
