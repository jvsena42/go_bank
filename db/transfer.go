package db

import (
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

func CreateTransfer(parameters dto.CreateTransferParameters) error {

	if !parameters.IsValid() {
		return errors.New("invalid parameters")
	}

	//TODO CHECK IF THE FromAccount is the ownwer and the ToAccount exist
	_, err := Db.Exec(createTransferQuery, parameters.FromAccountID, parameters.ToAccountID, parameters.Amount)

	return err
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
