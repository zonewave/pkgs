package stringutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zonewave/pkgs/randutil"
)

func TestSpaceRemoveAll(t *testing.T) {
	require.Equal(t, "test", SpaceRemoveAll(" te st "))
}

func TestRandString(t *testing.T) {

	tests := []struct {
		name   string
		lenNum int
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "rand string",
			lenNum: 10,
			want:   "cUbYhiZzKa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandString(tt.lenNum, randutil.NewRand(0)); got != tt.want {
				t.Errorf("RandString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpaceRemove(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "remove 3 space",
			args: args{
				s: "  te st ",
				n: 3,
			},
			want: "test ",
		},
		{
			name: "remove all space",
			args: args{
				s: "  te st ",
				n: -1,
			},
			want: "test",
		},
		{
			name: "remove all space",
			args: args{
				s: "  te st ",
				n: 0,
			},
			want: "  te st ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpaceRemove(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("SpaceRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
