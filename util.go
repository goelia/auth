package auth

import (
	"math/rand"
	"time"
)

// RandNum return length 6
func RandNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100000)
}
