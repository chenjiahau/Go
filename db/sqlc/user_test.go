package db

import (
	"context"
	"testing"

	"example.com/project/util"
	"github.com/stretchr/testify/require"
)

func CreateUser(t *testing.T) Users {
	arg := CreateUserParams{
		Username: util.GenerateRandomString(10),
		Password: util.GenerateRandomString(10),
		FullName: util.GenerateRandomString(10),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func DeleteUser(t *testing.T, user1 Users) {
	user2, err := testQueries.Delete(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
}

func DeleteAllUser(t *testing.T) {
	users, err := testQueries.DeleteAll(context.Background())

	require.NoError(t, err)
	require.Len(t, users, 10)
}

func TestCreateUser(t *testing.T) {
	user1 := CreateUser(t)
	DeleteUser(t, user1)
}

func TestGetUser(t *testing.T) {
	user1 := CreateUser(t)
	arg := GetUserParams{
		Username: user1.Username,
		Password: user1.Password,
	}
	user2, err := testQueries.GetUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)

	DeleteUser(t, user1)
}

func TestGetUserById(t *testing.T) {
	user1 := CreateUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)

	DeleteUser(t, user1)
}

func TestGetAllUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateUser(t)
	}

	users, err := testQueries.GetUsers(context.Background())

	require.NoError(t, err)
	require.Len(t, users, 10)

	DeleteAllUser(t)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateUser(t)

	arg := UpdateParams{
		ID: user1.ID,
		Username: util.GenerateRandomString(10),
		Password: util.GenerateRandomString(10),
		FullName: util.GenerateRandomString(10),
	}

	user2, err := testQueries.Update(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.FullName, user2.FullName)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)

	DeleteUser(t, user2)
}

func TestDeleteUser(t *testing.T) {
	user := CreateUser(t)
	DeleteUser(t, user)
}

func TestDeleteAllUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateUser(t)
	}

	DeleteAllUser(t)
}

