package randutil

import (
	"math/rand"
	"time"
)

// DefaultRand return a default rand
func DefaultRand() *rand.Rand {
	return NewRand(time.Now().UnixNano())
}

// NewRand return a new rand by seed
func NewRand(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}
