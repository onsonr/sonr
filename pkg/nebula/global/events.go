package global

type Event string

const (
	// Auth
	AuthLoginSuccess    Event = "login_success"
	AuthLogoutSuccess   Event = "logout_success"
	AuthRegisterSuccess Event = "register_success"
	AuthRegisterFailure Event = "register_failure"

	WalletTransactionConfirmed Event = "transaction_confirmed"
)
