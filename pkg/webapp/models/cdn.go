package models

import "fmt"

// ╭───────────────────────────────────────────────────────────╮
// │                  Avatar Image Components                  │
// ╰───────────────────────────────────────────────────────────╯

type Avatar string

const (
	Avatar0xDesigner    Avatar = "0xdesigner.jpg"
	AvatarAccountless   Avatar = "accountless.jpg"
	AvatarAlexRecouso   Avatar = "alexrecouso.jpg"
	AvatarChjango       Avatar = "chjango.jpg"
	AvatarGwart         Avatar = "gwart.jpg"
	AvatarHTMXOrg       Avatar = "htmx_org.jpg"
	AvatarJelenaNoble   Avatar = "jelena_noble.jpg"
	AvatarSonr          Avatar = "sonr.svg"
	AvatarTanishqXYZ    Avatar = "tanishqxyz.jpg"
	AvatarUnusualWhales Avatar = "unusual_whales.png"
	AvatarWinnieLaux    Avatar = "winnielaux_.jpg"
)

func (a Avatar) Src() string {
	return fmt.Sprintf("https://cdn.sonr.id/img/avatars/%s", string(a))
}

// ╭───────────────────────────────────────────────────────────╮
// │                  General Illustrations                    │
// ╰───────────────────────────────────────────────────────────╯

type Illustration string

const (
	BlockchainExplorer   Illustration = "blockchain-explorer"
	BlockchainStructure  Illustration = "blockchain-structure"
	CrossChainBridge     Illustration = "cross-chain-bridge"
	CrossChainTransfer   Illustration = "cross-chain-transfer"
	CryptoAirdrop        Illustration = "crypto-airdrop"
	CryptoCard           Illustration = "crypto-card"
	CryptoExchange       Illustration = "crypto-exchange"
	CryptoMining         Illustration = "crypto-mining"
	CryptoPayments       Illustration = "crypto-payments"
	CryptoSecurity       Illustration = "crypto-security"
	CryptoStaking        Illustration = "crypto-staking"
	CryptoYield          Illustration = "crypto-yield"
	CurrencyConversion   Illustration = "currency-conversion"
	DecentralizedNetwork Illustration = "decentralized-network"
	DecentralizedWebNode Illustration = "decentralized-web-node"
	DefiDashboard        Illustration = "defi-dashboard"
	GovernanceToken      Illustration = "governance-token"
	HardwareWallet       Illustration = "hardware-wallet"
	InitialCoinOffering  Illustration = "initial-coin-offering"
	LiquidityPool        Illustration = "liquidity-pool"
	MarketAnalysis       Illustration = "market-analysis"
	MarketVolatility     Illustration = "market-volatility"
	MultiCoinWallet      Illustration = "multi-coin-wallet"
	NetworkLatency       Illustration = "network-latency"
	PortfolioBalance     Illustration = "portfolio-balance"
	PrivateKey           Illustration = "private-key"
	ProofOfStake         Illustration = "proof-of-stake"
	TokenFractioning     Illustration = "token-fractioning"
	TokenMinting         Illustration = "token-minting"
	TokenSwap            Illustration = "token-swap"
)

func (i Illustration) Src() string {
	return fmt.Sprintf("https://cdn.sonr.id/img/illustrations/%s.svg", string(i))
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Empty State Illustrations                │
// ╰───────────────────────────────────────────────────────────╯

type ErrorIllustration string

const (
	ErrorIllustration404                ErrorIllustration = "404"
	ErrorIllustrationEmptyCart          ErrorIllustration = "empty-cart"
	ErrorIllustrationEmptyMessages      ErrorIllustration = "empty-messages"
	ErrorIllustrationEmptyNotifications ErrorIllustration = "empty-notifications"
	ErrorIllustrationEmptySearch        ErrorIllustration = "empty-search"
	ErrorIllustrationFileNotFound       ErrorIllustration = "file-not-found"
	ErrorIllustratinInternalError       ErrorIllustration = "internal-error"
	ErrorIllustrationNoConnection       ErrorIllustration = "no-connection"
	ErrorIllustrationOrderNotFound      ErrorIllustration = "order-not-found"
	ErrorIllustrationPaymentFailed      ErrorIllustration = "payment-failed"
)

func (e ErrorIllustration) SrcDefault() string {
	return fmt.Sprintf("https://cdn.sonr.id/img/empty-state/default/%s.svg", string(e))
}

func (e ErrorIllustration) SrcMono() string {
	return fmt.Sprintf("https://cdn.sonr.id/img/empty-state/mono/%s.svg", string(e))
}

func (e ErrorIllustration) SrcRetro() string {
	return fmt.Sprintf("https://cdn.sonr.id/img/empty-state/retro/%s.svg", string(e))
}