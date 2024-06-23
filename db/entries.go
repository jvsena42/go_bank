package db

import "github.com/jvsena42/go_bank/dto"

const createEntryQuery = `
INSERT INTO entries (
	account_id,
	amount
) VALUES (
	$1, $2
);
`

const listEntries = `
SELECT id, account_id, amount, created_at FROM entries 
WHERE account_id = $1
ORDER BY created_at;
`

func CreateEntry(parameters dto.CreateEntryParamets) error {
	_, err := Db.Exec(createEntryQuery, parameters.AccountId, parameters.Amount)

	return err
}

func ListEntries(accountId int64) ([]Entry, error) {
	rows, err := Db.Query(listEntries, accountId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(&entry.ID, &entry.AccountID, &entry.Amount, &entry.CreatedAt)

		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, err
}
