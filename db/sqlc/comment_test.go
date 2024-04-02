package db

import (
	"context"
	"database/sql"
	"testing"

	"example.com/project/util"
	"github.com/stretchr/testify/require"
)

func CreateComment(t *testing.T) (Users, EntryTypes, Entries ,Comments) {
	user := CreateUser(t)
	entryType, entry := CreateEntry(t)

	arg := CreateCommentParams{
		EntryID: entry.ID,
		UserID:  user.ID,
	}

	comment, err := testQueries.CreateComment(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.EntryID, comment.EntryID)
	require.Equal(t, arg.UserID, comment.UserID)
	require.Equal(t, arg.Content, comment.Content)

	require.NotZero(t, comment.ID)
	require.NotZero(t, comment.CreatedAt)

	return user, entryType, entry, comment
}

func DeleteComment(t *testing.T, user1 Users, entryType1 EntryTypes, entry1 Entries, comment1 Comments) {
	comment2, err := testQueries.DeleteComment(context.Background(), comment1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.EntryID, comment2.EntryID)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.Equal(t, comment1.Content, comment2.Content)

	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.CreatedAt, comment2.CreatedAt)

	DeleteEntry(t, entryType1, entry1)
	DeleteUser(t, user1)
}

func DeleteAllComments(t *testing.T) {
	comments, err := testQueries.DeleteComments(context.Background())

	require.NoError(t, err)
	require.Len(t, comments, 10)

	entries, err := testQueries.DeleteAllEntries(context.Background())

	require.NoError(t, err)
	require.Len(t, entries, 10)

	entryTypes, err := testQueries.DeleteAllEntryTypes(context.Background())

	require.NoError(t, err)
	require.Len(t, entryTypes, 10)

	users, err := testQueries.DeleteAll(context.Background())

	require.NoError(t, err)
	require.Len(t, users, 10)
}

func TestCreateComment(t *testing.T) {
	user1, entryType1, entry1, comment1 := CreateComment(t)
	DeleteComment(t, user1, entryType1, entry1, comment1)
}

func TestGetComment(t *testing.T) {
	user1, entryType1, entry1, comment1 := CreateComment(t)
	comment2, err := testQueries.GetComment(context.Background(), comment1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.EntryID, comment2.EntryID)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.Equal(t, comment1.Content, comment2.Content)

	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.CreatedAt, comment2.CreatedAt)

	DeleteComment(t, user1, entryType1, entry1, comment1)
}

func TestGetCommentById(t *testing.T) {
	user1, entryType1, entry1, comment1 := CreateComment(t)
	comment2, err := testQueries.GetCommentById(context.Background(), comment1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.EntryID, comment2.EntryID)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.Equal(t, comment1.Content, comment2.Content)

	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.CreatedAt, comment2.CreatedAt)

	DeleteComment(t, user1, entryType1, entry1, comment1)
}

func TestGetAllComments(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateComment(t)
	}

	comments, err := testQueries.GetComments(context.Background())

	require.NoError(t, err)
	require.Len(t, comments, 10)

  DeleteAllComments(t)
}

func TestUpdateComment(t *testing.T) {
	user1, entryType1, entry1, comment1 := CreateComment(t)

	arg := UpdateCommentParams{
		ID:      comment1.ID,
		Content: sql.NullString{String: util.GenerateRandomString(100), Valid: true},
	}

	comment2, err := testQueries.UpdateComment(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.EntryID, comment2.EntryID)
	require.Equal(t, comment1.UserID, comment2.UserID)

	require.Equal(t, arg.Content, comment2.Content)

	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.CreatedAt, comment2.CreatedAt)

	DeleteComment(t, user1, entryType1, entry1, comment2)
}

func TestDeleteComment(t *testing.T) {
	user1, entryType1, entry1, comment1 := CreateComment(t)
	DeleteComment(t, user1, entryType1, entry1, comment1)
}

func TestDeleteAllComments(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateComment(t)
	}

	DeleteAllComments(t)
}