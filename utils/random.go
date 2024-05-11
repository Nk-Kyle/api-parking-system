package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
