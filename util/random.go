package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefgijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) string {
	var stringBuilder strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		character := alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(character)
	}

	return stringBuilder.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 100000)
}

func RandomCurrency() string {
	currencies := []string{BTC, USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
