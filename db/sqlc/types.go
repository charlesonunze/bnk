package db

type (
	// TransferTxParams contains the input parameters of the transfer transaction
	TransferTxParams struct {
		FromAccountID int64 `json:"from_account_id"`
		ToAccountID   int64 `json:"to_account_id"`
		Amount        int64 `json:"amount"`
	}

	// TransferTxResult is the result of the transfer transaction
	TransferTxResult struct {
		Transfer    Transfer `json:"transfer"`
		FromEntry   Entry    `json:"from_entry"`
		ToEntry     Entry    `json:"to_entry"`
		FromAccount Account  `json:"from_account"`
		ToAccount   Account  `json:"to_account"`
	}
)
