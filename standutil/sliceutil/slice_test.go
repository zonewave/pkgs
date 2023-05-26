package sliceutil

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func TestItemInSlice(t *testing.T) {
	type args[T comparable] struct {
		item T
		list []T
	}
	type testNode[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	testStrings := []testNode[string]{
		{
			name: "string",
			args: args[string]{
				item: "222",
				list: []string{"233", "222", "555"},
			},
			want: true,
		},
		{
			name: "string not int",
			args: args[string]{
				item: "2227",
				list: []string{"233", "222", "555"},
			},
			want: false,
		},
	}
	testInts := []testNode[int]{
		{
			name: "int",
			args: args[int]{
				item: 222,
				list: []int{233, 222, 555},
			},
			want: true,
		},
	}
	for _, tt := range testStrings {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.item, tt.args.list); got != tt.want {
				t.Errorf("SliceContain() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.item, tt.args.list); got != tt.want {
				t.Errorf("SliceContain() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestGenerateSequences(t *testing.T) {
	type args[T constraints.Integer] struct {
		start T
		end   T
		step  T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "test",
			args: args[int]{
				start: 0,
				end:   5,
				step:  1,
			},
			want: []int{0, 1, 2, 3, 4},
		}, {
			name: "0 step",
			args: args[int]{
				start: 1,
				end:   5,
				step:  0,
			},
			want: []int{},
		},
		{
			name: "step>0,start > end",
			args: args[int]{
				start: 5,
				end:   1,
				step:  1,
			},
			want: []int{},
		},
		{
			name: "step<0,start < end",
			args: args[int]{
				start: 1,
				end:   5,
				step:  -1,
			},
			want: []int{},
		},
		{
			name: "step<0,start > end",
			args: args[int]{
				start: 5,
				end:   -1,
				step:  -1,
			},
			want: []int{5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSequences(tt.args.start, tt.args.end, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSequences() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestIterFn(t *testing.T) {

	// ok
	s := GenerateSequences[int](1, 5, 1)
	mul := 1
	fn := func(index int, item int) bool {
		mul *= item
		return true
	}

	s2 := IterFn(s, fn)
	require.Equal(t, mul, 24)
	require.Equal(t, s2, []int{1, 2, 3, 4})

	mul = 1
	IterFn(s, func(index int, item int) bool {
		if index == 2 {
			return false
		}
		mul *= item
		return true
	})
	require.Equal(t, mul, 2)

}

func TestMap(t *testing.T) {
	type args[T any, K any] struct {
		arr []T
		fn  func(item T) K
	}
	type testCase[T any, K any] struct {
		name string
		args args[T, K]
		want []K
	}
	tests := []testCase[int, string]{
		// TODO: Add test cases.
		{
			name: "test",
			args: args[int, string]{
				arr: GenerateSequences[int](1, 6, 1),
				fn: func(item int) string {
					return strconv.Itoa(item)
				},
			},
			want: []string{"1", "2", "3", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.arr, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotContain(t *testing.T) {
	type args[T comparable] struct {
		item T
		list []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "test",
			args: args[int]{
				item: 9,
				list: GenerateSequences[int](1, 6, 1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotContain(tt.args.item, tt.args.list); got != tt.want {
				t.Errorf("NotContain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args[T any] struct {
		arr []T
		fn  func(curItem, preCum T) T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "test",
			args: args[int]{
				arr: GenerateSequences[int](1, 6, 1),
				fn: func(curItem, preCum int) int {
					return curItem + preCum
				},
			},
			want: []int{1, 3, 6, 10, 15},
		},
		{
			name: "empty",
			args: args[int]{
				arr: nil,
				fn: func(curItem, preCum int) int {
					return curItem + preCum
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.arr, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args[T comparable, K interface {
		constraints.Integer | constraints.Float
	}] struct {
		list  []T
		value func(item T) K
	}
	type testCase[T comparable, K interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T, K]
		want K
	}
	tests := []testCase[int, int]{
		// TODO: Add test cases.
		{
			name: "test",
			args: args[int, int]{
				list:  GenerateSequences[int](1, 6, 1),
				value: func(item int) int { return item },
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.list, tt.args.value); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[T any] struct {
		arr []T
		fn  func(item T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test",
			args: args[int]{
				arr: GenerateSequences[int](1, 6, 1),
				fn: func(item int) bool {
					return item%2 == 0
				},
			},
			want: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.arr, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	type args[T any, K comparable] struct {
		slice []T
		id    func(T) K
	}
	type testCase[T any, K comparable] struct {
		name string
		args args[T, K]
		want map[K][]T
	}
	tests := []testCase[int, int]{
		// TODO: Add test cases.
		{
			name: "orderby Odd",
			args: args[int, int]{

				slice: GenerateSequences[int](1, 6, 1),
				id: func(item int) int {
					return item % 2
				},
			},
			want: map[int][]int{
				0: {2, 4},
				1: {1, 3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.slice, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
