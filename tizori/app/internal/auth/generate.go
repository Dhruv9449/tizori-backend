package auth

import (
	"math/rand"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

// GeneratePassword generates a random password of the given length
func GeneratePassword(length int, useLetters bool, useSpecial bool, useNum bool) string {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		if useLetters {
			b[i] = letterBytes[r1.Intn(len(letterBytes))]
		} else if useSpecial {
			b[i] = specialBytes[r1.Intn(len(specialBytes))]
		} else if useNum {
			b[i] = numBytes[r1.Intn(len(numBytes))]
		}
	}
	return string(b)
}
