package sliceutil

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestSequences(t *testing.T) {
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
			if got := Sequences(tt.args.start, tt.args.end, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sequences() = %v, want %v", got, tt.want)
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

				slice: Sequences[int](1, 6, 1),
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
