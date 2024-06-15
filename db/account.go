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

const listAccountsQuery = `
SELECT id, owner, balance, currency, created_at FROM accounts 
ORDER BY id
LIMIT $1
OFFSET $2;
`

const updateAccountQuery = `
UPDATE accounts 
SET balance = $1
WHERE id = $2;
`

func CreateAccount(parameters dto.CreateAccountParameters) error {
	_, err := Db.Exec(createAccountQuery, parameters.Owner, parameters.Balance, parameters.Currency)

	return err
}

func GetAccount(id int64) (*Account, error) {

	row := Db.QueryRow(getAccountQuery, id)

	account := &Account{}
	err := row.Scan(&account.ID, &account.Owner, &account.Balance, &account.Currency, &account.CreatedAt)

	if err != nil {
		return nil, err
	}

	return account, err
}

func ListAccounts() ([]Account, error) {
	rows, err := Db.Query(listAccountsQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		err := rows.Scan(&account.ID, &account.Owner, &account.Balance, &account.Currency, &account.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, err
}

func UpdateAccount(parameters dto.UpdateAccountParameters) error {
	_, err := Db.Exec(updateAccountQuery, parameters.Balance, parameters.ID)

	return err
}
