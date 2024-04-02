package db

import (
	"context"
	"database/sql"
	"testing"

	"example.com/project/util"
	"github.com/stretchr/testify/require"
)

func CreateEntry(t *testing.T) (EntryTypes, Entries) {
	entryType := CreateEntryType(t)
	arg := CreateEntryParams{
		EntryTypeID: entryType.ID,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.EntryTypeID, entry.EntryTypeID)
	require.Equal(t, arg.ReleaseVersion, entry.ReleaseVersion)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entryType, entry
}

func DeleteEntry(t *testing.T, entryType1 EntryTypes, entry1 Entries) {
	entry2, err := testQueries.DeleteEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.EntryTypeID, entry2.EntryTypeID)
	require.Equal(t, entry1.ReleaseVersion, entry2.ReleaseVersion)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)

	DeleteEntryType(t, entryType1)
}

func DeleteAllEntries(t *testing.T) {
	entries, err := testQueries.DeleteAllEntries(context.Background())

	require.NoError(t, err)
	require.Len(t, entries, 10)

	entryTypes, err := testQueries.DeleteAllEntryTypes(context.Background())

	require.NoError(t, err)
	require.Len(t, entryTypes, 10)
}

func TestCreateEntry(t *testing.T) {
	entryType1, entry1 := CreateEntry(t)
	DeleteEntry(t, entryType1, entry1)
}

func TestGetEntry(t *testing.T) {
	entryType1, entry1 := CreateEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.EntryTypeID, entry2.EntryTypeID)
	require.Equal(t, entry1.ReleaseVersion, entry2.ReleaseVersion)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)

	DeleteEntry(t, entryType1, entry1)
}

func TestGetEntryById(t *testing.T) {
	entryType1, entry1 := CreateEntry(t)
	entry2, err := testQueries.GetEntryById(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.EntryTypeID, entry2.EntryTypeID)
	require.Equal(t, entry1.ReleaseVersion, entry2.ReleaseVersion)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)

	DeleteEntry(t, entryType1, entry1)
}

func TestGetAllEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateEntry(t)
	}

	entries, err := testQueries.GetEntries(context.Background())

	require.NoError(t, err)
	require.Len(t, entries, 10)

	DeleteAllEntries(t)
}

func TestUpdateEntry(t *testing.T) {
	entryType1, entry1 := CreateEntry(t)

	arg := UpdateEntryParams{
		ID:              entry1.ID,
		EntryTypeID:     entryType1.ID,
		ReleaseVersion:  sql.NullString{String: util.GenerateRandomString(10), Valid: true},
		Priority:        sql.NullInt32{Int32: 2, Valid: true},
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, arg.EntryTypeID, entry2.EntryTypeID)
	require.Equal(t, arg.ReleaseVersion, entry2.ReleaseVersion)
	require.Equal(t, arg.Priority, entry2.Priority)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)

	DeleteEntry(t, entryType1, entry2)
}

func TestDeleteEntry(t *testing.T) {
	entryType1, entry1 := CreateEntry(t)
	DeleteEntry(t, entryType1, entry1)
}

func TestDeleteAllEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateEntry(t)
	}

	DeleteAllEntries(t)
}