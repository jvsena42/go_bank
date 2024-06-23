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

func CreateTransfer(parameters dto.CreateTransferParameters) error {

	if !parameters.IsValid() {
		return errors.New("invalid parameters")
	}

	//TODO CHECK IF THE FromAccount is the ownwer and the ToAccount exist
	_, err := Db.Exec(createTransferQuery, parameters.FromAccountID, parameters.ToAccountID, parameters.Amount)

	return err
}
