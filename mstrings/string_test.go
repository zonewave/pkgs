package mstrings

import "testing"

func TestSpaceRemoveAll(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"no space",
			args{
				"test",
			},
			"test",
		},
		{
			"hasspace",
			args{
				" te st ",
			},
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpaceRemoveAll(tt.args.s); got != tt.want {
				t.Errorf("SpaceRemoveAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
