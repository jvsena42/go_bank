package db

import (
	"github.com/jvsena42/go_bank/dto"
)

const createAccountQuery = `
INSERT INTO accounts (
	owner,
	balance,
	currency
) VALUES (
	$1, $2, $3
)
`

func CreateAccount(parameters dto.CreateAccountParameters) error {

	_, err := Db.Exec(createAccountQuery, parameters.Owner, parameters.Balance, parameters.Currency)

	return err
}
