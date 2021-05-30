package server

import (
	"math/rand"
	"time"
)

func randomPort(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
