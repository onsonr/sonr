package sections

import "fmt"

// ╭───────────────────────────────────────────────────────────╮
// │                  SVG CDN Illustrations                    │
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
