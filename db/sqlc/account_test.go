package db

import (
	"context"
	"github.com/m0nk-tnd/go-web-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomString(6),
		Balance:  util.RandomInt(1, 1000),
		Currency: util.RandomCurrency(),
	}
	return createAccountFunc(t, arg)
}

func createAccountFunc(t *testing.T, arg CreateAccountParams) Account {
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomString(6),
		Balance:  util.RandomInt(1, 1000),
		Currency: util.RandomCurrency(),
	}
	account := createAccountFunc(t, arg)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestQueries_GetAccount(t *testing.T) {
	expectedAccount := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), expectedAccount.ID)

	require.NoError(t, err)
	require.Equal(t, expectedAccount, account)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountParams{
		account.ID,
		util.RandomInt(0, 1000),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestQueries_ListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{
		5,
		5,
	}
	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
