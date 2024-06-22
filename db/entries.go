package db

import "github.com/jvsena42/go_bank/dto"

const createEntryQuery = `
INSERT INTO entries (
	account_id,
	amount,
) VALUES (
	$1, $2
);
`

func CreateEntry(parameters dto.CreateEntryParamets) error {
	_, err := Db.Exec(createEntryQuery, parameters.AccountId, parameters.Amount)

	return err
}
