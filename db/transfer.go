package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jvsena42/go_bank/dto"
)

const createTransferQuery = `
INSERT INTO transfers (
	from_account_id,
	to_account_id,
	amount
) VALUES (
	$1, $2, $3
);
`

const listTransfersQuery = `
	SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
	WHERE (from_account_id = $1 OR to_account_id = $1)
	ORDER BY created_at;
`

/*
Transfer 10 USD from account 1 to account 2
1 - Create a transfer record with amount 10
2 - Create an account entry for account 1 with amount -10
3 - Create an account entry for account 2 with amount 10
4 - Check if account 1 has enougth balance
4 - Subtract 10 from the  balance of account 1
5 - Add 10 to the balance of account 2
*/
func CreateTransfer(parameters dto.CreateTransferParameters, ctx context.Context) error {

	if !parameters.IsValid() {
		return errors.New("invalid parameters")
	}

	// Get a Tx for making transaction requests.
	tx, err := Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	//Create a transfer record with amount parameters.Amount
	_, err = tx.QueryContext(ctx, createTransferQuery, parameters.FromAccountID, parameters.ToAccountID, parameters.Amount)
	if err != nil {
		return err
	}

	//Create an account entry for FromAccountID 1 with amount -Amount
	_, err = tx.QueryContext(ctx, CreateEntryQuery, parameters.FromAccountID, -parameters.Amount)
	if err != nil {
		return err
	}

	//Create an account entry for ToAccountID 1 with amount Amount
	_, err = tx.QueryContext(ctx, CreateEntryQuery, parameters.ToAccountID, parameters.Amount)
	if err != nil {
		return err
	}

	//Get account 1
	account1 := &Account{}
	if err = tx.QueryRowContext(ctx, GetAccountQuery, parameters.FromAccountID).Scan(&account1); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("account not found")
		}
		return err
	}

	//check if account 1 has enougth balance
	if account1.Balance < parameters.Amount {
		return errors.New("not enougth balance")
	}

	//Update balance from account 1
	_, err = tx.QueryContext(ctx, updateAccountQuery, account1.Balance-parameters.Amount, parameters.FromAccountID)
	if err != nil {
		return err
	}

	//Get account 2
	account2 := &Account{}
	if err = tx.QueryRowContext(ctx, GetAccountQuery, parameters.FromAccountID).Scan(&account2); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("account not found")
		}
		return err
	}

	//Update balance from account 2
	_, err = tx.QueryContext(ctx, updateAccountQuery, account2.Balance+parameters.Amount, parameters.ToAccountID)
	if err != nil {
		return err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ListTransfers(accountId int64) ([]Transfer, error) {
	rows, err := Db.Query(listTransfersQuery, accountId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transfers []Transfer
	for rows.Next() {
		var transfer Transfer
		err := rows.Scan(&transfer.ID, &transfer.FromAccountID, &transfer.ToAccountID, &transfer.Amount, &transfer.CreatedAt)

		if err != nil {
			return nil, err
		}

		transfers = append(transfers, transfer)
	}

	return transfers, err
}
