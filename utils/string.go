package utils

import "math/rand"

var (
	DefaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~?<>!@#$%^&*({})+|")
)

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = DefaultLetters[rand.Intn(len(DefaultLetters))]
	}

	return string(b)
}
