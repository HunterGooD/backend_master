package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/HunterGooD/backend_master/internal/utils"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accCreated := createRandomAccount(t)
	accGeting, err := testQueries.GetAccount(context.Background(), accCreated.ID)

	require.NoError(t, err)
	require.NotEmpty(t, accGeting)

	require.Equal(t, accCreated.ID, accGeting.ID)
	require.Equal(t, accCreated.Owner, accGeting.Owner)
	require.Equal(t, accCreated.Balance, accGeting.Balance)
	require.Equal(t, accCreated.Currency, accGeting.Currency)
	// отметки времени различаются не более чем на дельта, тут секунда
	require.WithinDuration(t, accCreated.CreatedAt, accGeting.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accCreated := createRandomAccount(t)

	updateArgs := UpdateAccountParams{
		ID:      accCreated.ID,
		Balance: utils.RandomMoney(),
	}

	updateAccount, err := testQueries.UpdateAccount(context.Background(), updateArgs)

	require.NoError(t, err)
	require.NotEmpty(t, updateAccount)

	require.Equal(t, accCreated.ID, updateAccount.ID)
	require.Equal(t, accCreated.Owner, updateAccount.Owner)
	require.Equal(t, updateArgs.Balance, updateAccount.Balance)
	require.Equal(t, accCreated.Currency, updateAccount.Currency)
	require.WithinDuration(t, accCreated.CreatedAt, updateAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	accCreated := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), accCreated.ID)

	require.NoError(t, err)
	accGeting, err := testQueries.GetAccount(context.Background(), accCreated.ID)
	require.Error(t, err)
	// проверка та ли ошибка выскочила а не какая нибудь другая
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accGeting)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), args)

	// нет ошибки и структура не пустая
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	// соответсвует тем данным которые создаются
	require.Equal(t, args.Owner, acc.Owner)
	require.Equal(t, args.Balance, acc.Balance)
	require.Equal(t, args.Currency, acc.Currency)

	// не равны нулю
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}
