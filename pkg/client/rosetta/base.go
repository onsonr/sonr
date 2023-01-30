package rosetta

// Client is the interface that wraps the basic methods of a Rosetta client.
// This interface is implemented in the Fine-Tuned Accounts in sonrhq/pkg/crypto/wallet,
// (i.e. BTC, ETH, Cosmos, etc.)

// Client is the interface that wraps the basic methods of a Rosetta client.
type Client interface {
	// GetAccountBalance returns the balance of the given account.
	GetAccountBalance(accountID string) (uint64, error)

	// GetAccountTransactions returns the transactions of the given account.
	GetAccountTransactions(accountID string) ([]string, error)

	// GetTransaction returns the transaction with the given hash.
	GetTransaction(txHash string) (string, error)

	// GetBlock returns the block with the given hash.
	GetBlock(blockHash string) (string, error)

	// GetBlockTransactions returns the transactions of the given block.
	GetBlockTransactions(blockHash string) ([]string, error)

	// GetBlockTransaction returns the transaction with the given hash in the given block.
	GetBlockTransaction(blockHash, txHash string) (string, error)

	// GetNetworkStatus returns the network status.
	GetNetworkStatus() (string, error)

	// SendTransaction sends the given transaction.
	SendTransaction(tx string) (string, error)
}
