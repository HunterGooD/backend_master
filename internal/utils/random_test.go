package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	a := RandomOwner()
	str := RandomString(10)

	require.NotEmpty(t, a)
	require.Len(t, a, 6)

	require.NotEmpty(t, str)
	require.Len(t, str, 10)
}

func TestRandomMoney(t *testing.T) {
	money := RandomMoney()

	require.True(t, money <= 2000 && money >= 0)
}

func TestRandomCurency(t *testing.T) {
	currencies := []string{"EUR", "USD", "RUS"}

	cr := RandomCurrency()
	inc := false

	for _, el := range currencies {
		if el == cr {
			inc = true
			break
		}
	}
	require.True(t, inc)
}
