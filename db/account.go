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
);
`

const getAccountQuery = `
SELECT id, owner, balance, currency, created_at FROM accounts 
WHERE id = $1 LIMIT 1;
`

func CreateAccount(parameters dto.CreateAccountParameters) error {
	_, err := Db.Exec(createAccountQuery, parameters.Owner, parameters.Balance, parameters.Currency)

	return err
}

func GetAccount(id int64) error {
	_, err := Db.Exec(getAccountQuery, id)

	return err
}
