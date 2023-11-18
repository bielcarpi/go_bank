package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max (inclusive).
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	// We could also use strings.Builder
	// Make a byte slice of length n
	s := make([]byte, n)

	// Fill the byte slice with random letters
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	// Return the string representation of the byte slice
	return string(s)
}

func RandomOwner() string {
	return RandomString(8)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := GetCurrencies()
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
