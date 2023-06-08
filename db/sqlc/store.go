package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx execute a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams contain the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// struct with empty object to store key transaction
//var txKey = struct{}{}

// TransferTx performs a money transfer from one account to the other
// It creates a transfer record, add account entries, and update accounts balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// txName := ctx.Value(txKey)
		// fmt.Println(txName, "create transfer")

		//transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		})
		if err != nil {
			return err
		}

		// add from account entries
		//fmt.Println(txName, "create entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})
		if err != nil {
			return err
		}

		// add to account entries
		//fmt.Println(txName, "create entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})
		if err != nil {
			return err
		}

		//TODO: update accounts balance
		//move money from - to account
		//fmt.Println(txName, "get account 1")
		// account1, err := q.GetAccountForUpdate(ctx, args.FromAccountID)
		// if err != nil {
		// 	return err
		// }

		//fmt.Println(txName, "update account 1")
		// result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
		// 	ID:      args.FromAccountID,
		// 	Balance: account1.Balance - args.Amount,
		// })
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     args.FromAccountID,
			Amount: -args.Amount,
		})
		if err != nil {
			return err
		}
		//move money to - from account
		//fmt.Println(txName, "get account 2")
		// account2, err := q.GetAccountForUpdate(ctx, args.ToAccountID)
		// if err != nil {
		// 	return err
		// }

		//fmt.Println(txName, "update account 2")
		// result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
		// 	ID:      args.ToAccountID,
		// 	Balance: account2.Balance + args.Amount,
		// })
		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     args.ToAccountID,
			Amount: args.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
