package routes

const (
	// Accounts endpoints
	kCreateAccountEndpoint     = "/sonr/highway/accounts/create/:coinType/:name" // Authenticated
	kGetAccountEndpoint        = "/sonr/highway/accounts/:did" // Public
	kListAccountsEndpoint      = "/sonr/highway/accounts/:coinType" // Authenticated
	kSignWithAccountEndpoint   = "/sonr/highway/accounts/:did/sign" // Authenticated
	kVerifyWithAccountEndpoint = "/sonr/highway/accounts/:did/verify" // Public

	// Maibox endpoints
	kGetMailboxMessagesEndpoint    = "/sonr/highway/mailbox/:did" // Authenticated
	kOpenMailboxMessageEndpoint    = "/sonr/highway/mailbox/:did/:messageId/read" // Authenticated
	kSendMailboxMessageEndpoint    = "/sonr/highway/mailbox/:from/:to/send" // Authenticated
	kDeleteMailboxMessageEndpoint  = "/sonr/highway/mailbox/:did/:messageId/delete" // Authenticated
	kPublishChannelMessageEndpoint = "/sonr/highway/mailbox/channels/:channelId/publish" // Authenticated
	kSubscribeChannelEndpoint      = "/sonr/highway/mailbox/channels/:channelId/subscribe" // Authenticated

	// Services endpoints
	kGetCredCreationOptionsEndpoint = "/sonr/highway/services/:origin/register/start/:alias" // Public
	kGetCredAssertionOptionsEndpoint = "/sonr/highway/services/:origin/login/start/:alias" // Public
	kRegisterCredForClaimsEndpoint  = "/sonr/highway/services/:origin/register/finish/:alias" // Public
	kVerifyCredForAccountEndpoint   = "/sonr/highway/services/:origin/login/finish/:alias" // Public

	// TX endpoints
	kBroadcastSonrTxEndpoint = "/sonr/highway/tx/snr/broadcast" // Public
	kGetSonrTxEndpoint       = "/sonr/highway/tx/snr/:txHash" // Public
)
