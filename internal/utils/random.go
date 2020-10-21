package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

// RandomInt генерирует случайное число в промежутке
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString генерирует рандомную строку только из символов поэтому random.Read не работает
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner случайное имя пользователя
func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney случайное кол-во монет
func RandomMoney() int64 {
	return RandomInt(0, 2000)
}

// RandomCurrency случайная валюта
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "RUS"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
