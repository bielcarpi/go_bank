package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"go_bank/util"
	"testing"
)

func createRandomTransfer(t *testing.T) Transfer {
	// Create two new accounts
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// Create a new transfer
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	// Check if the transfer is created successfully
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	// Create a new transfer and retrieve it
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	// Check if the transfer is retrieved successfully
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.CreatedAt, transfer2.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	// Create 10 new transfers
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	// Retrieve all transfers from account 1 to account 2
	arg := ListTransfersParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Limit:         5,
		Offset:        0,
	}

	_, err := testQueries.ListTransfers(context.Background(), arg)

	// Check if the transfers are retrieved successfully
	require.NoError(t, err)
}
