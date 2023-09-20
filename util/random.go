package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var MyRand *rand.Rand

func init() {
	MyRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt returns random int between min and max
func RandomInt(min, max int64) int64 {
	return min + MyRand.Int63n(max-min+1)
}

// RandomString returns random string of specified length
func RandomString(length int) string {
	var sb strings.Builder
	AbcLen := len(alphabet)

	for i := 0; i < length; i++ {
		ch := alphabet[MyRand.Intn(AbcLen)]
		sb.WriteByte(ch)
	}

	return sb.String()
}

// RandomCurrency returns random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "RUB"}
	return currencies[MyRand.Intn(len(currencies))]
}
