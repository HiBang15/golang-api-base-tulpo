package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	ascii_letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ascii_lowercase = "abcdefghijklmnopqrstuvwxyz"
	ascii_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"


)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max - min + 1)
}

func RandomInt32(min, max int32) int32 {
	return min + rand.Int31n(max - min + 1)
}

func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	l := len(ascii_letters)

	for i := 0; i < n; i++ {
		c := ascii_letters[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomStringFromList(list ...string) string {
	l := len(list)
	if l == 0 {
		return ""
	}
	return list[rand.Intn(l)]
}