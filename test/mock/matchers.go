package mock

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"reflect"

	"go.uber.org/mock/gomock"
)

type eqMatcher[T any] struct {
	x T
}

func (e eqMatcher[T]) Matches(src any) bool {
	// In case, some value is nil
	x, ok := src.(T)
	if !ok {
		return false
	}
	return reflect.DeepEqual(e.x, x)
}

func (e eqMatcher[T]) String() string {
	return fmt.Sprintf("is equal to %v (%T)", e.x, e.x)
}

type nilMatcher[T any] struct{}

func (nilMatcher[T]) Matches(src any) bool {
	if src == nil {
		return true
	}
	x, ok := src.(T)
	if !ok {
		return false
	}
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}

	return false
}

func (nilMatcher[T]) String() string {
	return "is nil"
}

type inAnyOrderMatcher[T any] struct {
	x []T
}

func (e inAnyOrderMatcher[T]) Matches(x any) bool {
	given, ok := x.([]T)
	if !ok {
		return false
	}
	if len(given) != len(e.x) {
		return false
	}
	givenLen := len(given)
	wantLen := len(e.x)

	usedFromGiven := make([]bool, givenLen)
	foundFromWanted := make([]bool, wantLen)
	for i := 0; i < wantLen; i++ {
		wantedMatcher := Eq(e.x[i])
		for j := 0; j < givenLen; j++ {
			if usedFromGiven[j] {
				continue
			}
			if wantedMatcher.Matches(given[j]) {
				foundFromWanted[i] = true
				usedFromGiven[j] = true
				break
			}
		}
	}

	falseFind := func(_ int, b bool) bool {
		return !b
	}
	missingFromWanted := slice.CountBy(foundFromWanted, falseFind)
	extraInGiven := slice.CountBy(usedFromGiven, falseFind)
	return extraInGiven == 0 && missingFromWanted == 0
}

func (e inAnyOrderMatcher[T]) String() string {
	return fmt.Sprintf("has the same elements as %v", e.x)
}

// CondMatcher is a Matcher that returns true for match function
type CondMatcher[T any] struct {
	match  func(val T) bool
	fields map[string]any
	format func(fields map[string]any)
}

func (e *CondMatcher[T]) Matches(src any) bool {
	x, ok := src.(T)
	if !ok {
		return false
	}
	return e.match(x)
}
func (e *CondMatcher[T]) String() string {
	bs, _ := json.Marshal(e.fields)
	return "CondMatcher:" + string(bs)
}

// WithField add field
func WithField[T any](key string, val any) func(matcher *CondMatcher[T]) {
	return func(matcher *CondMatcher[T]) {
		matcher.fields[key] = val
	}
}

// WithFormat add format
func WithFormat[T any](fn func(fields map[string]any)) func(matcher *CondMatcher[T]) {
	return func(matcher *CondMatcher[T]) {
		matcher.format = fn
	}
}

// WithFields add fields
func WithFields[T any](params ...any) func(matcher *CondMatcher[T]) {
	return func(matcher *CondMatcher[T]) {
		for i := 0; i < len(params); i += 2 {
			k := params[i].(string)
			v := params[i+1]
			matcher.fields[k] = v
		}
	}
}

// Eq is a Matcher that returns true when the argument is equal to the given value.
func Eq[T any](x T) gomock.Matcher {
	return eqMatcher[T]{x}
}

// Nil is a Matcher that returns true when the argument is nil.
func Nil[T any]() gomock.Matcher {
	return nilMatcher[T]{}
}

// Cond is a Matcher that returns true for match function
func Cond[T any](fn func(val T) bool, opts ...func(matcher *CondMatcher[T])) gomock.Matcher {
	ret := &CondMatcher[T]{match: fn}
	ret.fields = make(map[string]any)
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

// InAnyOrder is a Matcher that returns true for collections of the same elements ignoring the order.
//
// Example usage:
//
//	InAnyOrder([]int{1, 2, 3}).Matches([]int{1, 3, 2}) // returns true
//	InAnyOrder([]int{1, 2, 3}).Matches([]int{1, 2}) // returns false
func InAnyOrder[T any](x []T) gomock.Matcher {
	return inAnyOrderMatcher[T]{x: x}
}
