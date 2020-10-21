package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    "Влад",
		Balance:  1300,
		Currency: "RUB",
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
}
