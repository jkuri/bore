package server

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
)

func randomPort(min, max int) int {
	return rand.Intn(max-min) + min
}

func randID() string {
	b := make([]byte, 4)
	crand.Read(b)
	return fmt.Sprintf("%x", b)
}
