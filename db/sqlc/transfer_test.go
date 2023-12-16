package db

import (
	"context"
	"github.com/m0nk-tnd/go-web-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTransfer(t *testing.T, account1, account2 *Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomInt(1, 1000),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomInt(1, 1000),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestQueries_GetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	expectedTransfer := createRandomTransfer(t, &account1, &account2)
	transfer, err := testQueries.GetTransfer(context.Background(), expectedTransfer.ID)

	require.NoError(t, err)
	require.Equal(t, expectedTransfer, transfer)
}

func TestQueries_ListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	account3 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, &account1, &account2)
	}
	for i := 0; i < 3; i++ {
		createRandomTransfer(t, &account1, &account3)
	}

	arg := ListTransfersParams{
		account1.ID,
		account2.ID,
		5,
		5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

	arg.Offset = 0
	arg.Limit = 15
	transfers, err = testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 13)
}
