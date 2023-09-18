package idgenerator

import (
	"fmt"
	"sync"

	"golang.org/x/exp/constraints"
)

type LocalIdGenerator[T constraints.Integer] struct {
	id         T
	Start, End T
	mu         sync.Mutex
}

func NewLocalIdGenerator[T constraints.Integer](start, end T) *LocalIdGenerator[T] {
	return &LocalIdGenerator[T]{
		id:    start,
		Start: start,
		End:   end,
	}
}

func (l *LocalIdGenerator[T]) Next() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	ret := l.id
	l.id++
	if l.id > l.End {
		return l.End, fmt.Errorf("out of range,%d", l.id)
	}
	return ret, nil
}
func (l *LocalIdGenerator[T]) Nexts(sum int) []T {
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
