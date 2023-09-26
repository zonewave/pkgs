package idgenerator

import (
	"fmt"
	"sync"

	"golang.org/x/exp/constraints"
)

// LocalIDGenerator for generating id.
type LocalIDGenerator[T constraints.Integer] struct {
	id         T
	Start, End T
	mu         sync.Mutex
}

// NewLocalIDGenerator create a new LocalIDGenerator.
func NewLocalIDGenerator[T constraints.Integer](start, end T) *LocalIDGenerator[T] {
	return &LocalIDGenerator[T]{
		id:    start,
		Start: start,
		End:   end,
	}
}

// Next generate next id.
func (l *LocalIDGenerator[T]) Next() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	ret := l.id
	l.id++
	if l.id > l.End {
		return l.End, fmt.Errorf("out of range,%d", l.id)
	}
	return ret, nil
}

// Nexts batch generate next id.
func (l *LocalIDGenerator[T]) Nexts(sum int) []T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.id+T(sum) > l.End {
		panic("out of range")
	}
	ret := make([]T, sum)
	for i := 0; i < sum; i++ {
		ret[i] = l.id
		l.id++
	}
	return ret
}
