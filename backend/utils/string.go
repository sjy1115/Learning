package utils

import (
	"math/rand"
	"strings"
)

var (
	DefaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = DefaultLetters[rand.Intn(len(DefaultLetters))]
	}

	return string(b)
}

func EscapeLIKE(search string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(strings.ReplaceAll(search, `\`, `\\`),
				"_", `\_`),
			"%", `\%`),
		"'", `\'`)
}
