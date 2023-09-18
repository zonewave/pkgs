package mock

import (
	"encoding/json"

	"go.uber.org/mock/gomock"
)

type CondMatcher[T any] struct {
	match  func(val T) bool
	fields map[string]any
	format func(fields map[string]any)
}

func WithField[T any](key string, val any) func(matcher *CondMatcher[T]) {
	return func(matcher *CondMatcher[T]) {
		matcher.fields[key] = val
	}
}
func WithFormat[T any](fn func(fields map[string]any)) func(matcher *CondMatcher[T]) {
	return func(matcher *CondMatcher[T]) {
		matcher.format = fn
	}
}

func WithFields[T any](params ...any) func(matcher *CondMatcher[T]) {
	return func(matcher *CondMatcher[T]) {
		for i := 0; i < len(params); i += 2 {
			k := params[i].(string)
			v := params[i+1]
			matcher.fields[k] = v
		}
	}
}

func Cond[T any](fn func(val T) bool, opts ...func(matcher *CondMatcher[T])) gomock.Matcher {
	ret := &CondMatcher[T]{match: fn}
	ret.fields = make(map[string]any)
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}
func (e *CondMatcher[T]) Matches(x interface{}) bool {
	item := x.(T)
	return e.match(item)
}
func (e *CondMatcher[T]) String() string {
	bs, _ := json.Marshal(e.fields)
	return "CondMatcher:" + string(bs)
}
