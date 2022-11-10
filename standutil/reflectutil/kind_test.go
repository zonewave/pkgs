package reflectutil

import (
	"testing"
)

func TestIsStructPtr(t *testing.T) {

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "struct ptr",
			obj:  &struct{}{},
			want: true,
		},
		{
			name: "struct",
			obj:  struct{}{},
			want: false,
		},
		{
			name: "ptr",
			obj:  new(struct{}),
			want: true,
		},
		{
			name: "nil",
			obj:  nil,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStructPtr(tt.obj); got != tt.want {
				t.Errorf("IsStructPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
