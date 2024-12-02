package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jvsena42/go_bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

/*func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount()
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account1.CreatedAt, time.Second)

	require.NotZero(t, account2.ID)
	require.NotZero(t, account2.CreatedAt)
}*/

/*func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount()

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: account1.Balance,
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account1.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account1.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, account1.CreatedAt, time.Second)

	require.NotZero(t, updatedAccount.ID)
	require.NotZero(t, updatedAccount.CreatedAt)
}*/

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount()

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

/*func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount()
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}*/

func createRandomAccount() Account {
	arg := CreateAccountParams{
		Owner:    util.RandomName(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, _ := testQueries.CreateAccount(context.Background(), arg)

	return account
}
