package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"go_bank/util"
	"testing"
)

func createRandomEntry(t *testing.T) Entry {
	// Create a new account
	account := createRandomAccount(t)

	// Create a new entry
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    100,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func createEntries(t *testing.T, n int) []Entry {
	// Create a new account
	account := createRandomAccount(t)

	// Create n new entries
	entries := make([]Entry, n)
	for i := 0; i < n; i++ {
		arg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		}
		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
		require.Equal(t, arg.Amount, entry.Amount)
		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.CreatedAt)

		entries[i] = entry
	}

	return entries
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	// Create a new entry and retrieve it
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	// Check if the entry is retrieved successfully
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
}

func TestListEntries(t *testing.T) {
	// Create 10 new entries for the same account
	createdEntries := createEntries(t, 10)

	// Retrieve all entries for the account
	arg := ListEntriesParams{
		AccountID: createdEntries[0].AccountID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	// Check if the entries are retrieved successfully
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
