package shuffle

import (
	"crypto/rand"
	"math/big"
)

const (
	uppercases = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercases = "abcdefghijklmnopqrstuvwxyz"
	numbers    = "0123456789"
	symbols    = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	charset    = uppercases + lowercases + numbers + symbols
)

func Character(length int) string {
	var r string
	csl := big.NewInt(int64(len(charset)))

	var str = make([]byte, length)
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, csl)
		str[i] = charset[num.Int64()]
	}

	r = string(str)

	return r
}

func Number(length int) string {
	var r string

	str := make([]byte, length)
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		str[i] = numbers[n.Int64()]
	}

	r = string(str)

	return r
}
