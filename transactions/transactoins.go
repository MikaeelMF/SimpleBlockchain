package transactions

type Transaction struct {
}

// Write will write to UTXO database if the transaction is valid
func Write(tx Transaction) error

// GetBalance will read an account balance by reading UTXO database and returnes the balance
// and its set of Unspent transactions
func GetBalance(address string) (int, []*Transaction)

// Validates a set of transactions. Returnes an error if a transaction is invalid
// containing the faulty transaction.
func Validate(txs []*Transaction) error
