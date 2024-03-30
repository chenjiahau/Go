package db

import (
	"context"
	"testing"

	"example.com/project/util"
	"github.com/stretchr/testify/require"
)

func CreateEntryType(t *testing.T) EntryTypes {
	name := util.GenerateRandomString(10)
	entryType, err := testQueries.CreateEntryType(context.Background(), name)

	require.NoError(t, err)
	require.NotEmpty(t, entryType)

	require.Equal(t, name, entryType.Name)

	require.NotZero(t, entryType.ID)
	require.NotZero(t, entryType.CreatedAt)

	return entryType
}

func DeleteEntryType(t *testing.T, entryType1 EntryTypes) {
	entryType2, err := testQueries.DeleteEntryType(context.Background(), entryType1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entryType2)

	require.Equal(t, entryType1.Name, entryType2.Name)

	require.Equal(t, entryType1.ID, entryType2.ID)
	require.Equal(t, entryType1.CreatedAt, entryType2.CreatedAt)
}

func DeleteAllEntryType(t *testing.T) {
	entryTypes, err := testQueries.DeleteAllEntryTypes(context.Background())

	require.NoError(t, err)
	require.Len(t, entryTypes, 10)
}

func TestCreateEntryType(t *testing.T) {
	entryType := CreateEntryType(t)
	DeleteEntryType(t, entryType)
}

func TestGetEntryType(t *testing.T) {
	entryType := CreateEntryType(t)
	entryType2, err := testQueries.GetEntryType(context.Background(), entryType.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entryType2)

	require.Equal(t, entryType.Name, entryType2.Name)

	require.Equal(t, entryType.ID, entryType2.ID)
	require.Equal(t, entryType.CreatedAt, entryType2.CreatedAt)

	DeleteEntryType(t, entryType)
}

func TestGetEntryTypeById(t *testing.T) {
	entryType := CreateEntryType(t)
	entryType2, err := testQueries.GetEntryType(context.Background(), entryType.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entryType2)

	require.Equal(t, entryType.Name, entryType2.Name)

	require.Equal(t, entryType.ID, entryType2.ID)
	require.Equal(t, entryType.CreatedAt, entryType2.CreatedAt)

	DeleteEntryType(t, entryType)
}

func TestGetAllEntryTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateEntryType(t)
	}

	entryTypes, err := testQueries.GetEntryTypes(context.Background())

	require.NoError(t, err)
	require.Len(t, entryTypes, 10)

	for _, entryType := range entryTypes {
		require.NotEmpty(t, entryType)
	}

	DeleteAllEntryType(t)
}

func TestUpdateEntryType(t *testing.T) {
	entryType1 := CreateEntryType(t)

	arg := UpdateEntryTypeParams{
		ID:   entryType1.ID,
		Name: util.GenerateRandomString(10),
	}

	entryType2, err := testQueries.UpdateEntryType(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entryType2)

	require.Equal(t, arg.ID, entryType2.ID)
	require.Equal(t, arg.Name, entryType2.Name)

	DeleteEntryType(t, entryType2)
}

func TestDeleteEntryType(t *testing.T) {
	entryType := CreateEntryType(t)
	DeleteEntryType(t, entryType)
}

func TestDeleteAllEntryType(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateEntryType(t)
	}

	DeleteAllEntryType(t)
}