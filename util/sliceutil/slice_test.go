package sliceutil

import "testing"

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
			if got := SliceContain(tt.args.item, tt.args.list); got != tt.want {
				t.Errorf("SliceContain() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContain(tt.args.item, tt.args.list); got != tt.want {
				t.Errorf("SliceContain() = %v, want %v", got, tt.want)
			}
		})
	}

}
