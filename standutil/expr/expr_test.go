package expr

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestCondExpr(t *testing.T) {
	ret := CondExpr(true, 1, 2)
	require.Equal(t, 1, ret)
	ret = CondExpr(false, 1, 2)
	require.Equal(t, 2, ret)

}

func TestFirstOrDefault(t *testing.T) {

	type Foo struct {
		A int
	}
	type args struct {
		params       []*Foo
		defaultValue *Foo
	}
	type testCase struct {
		name string
		args args
		want *Foo
	}
	tests := []testCase{
		{
			name: "case1",
			args: args{
				params:       []*Foo{{A: 1}, {A: 2}},
				defaultValue: &Foo{A: 3},
			},
			want: &Foo{A: 1},
		},
		{
			name: "case2",
			args: args{
				defaultValue: &Foo{A: 3},
			},
			want: &Foo{A: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOrDefault(tt.args.params, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	require.Equal(t, 3, Add(1, 2))
	require.Equal(t, 0.8, Add(0.5, 0.3))
	require.Equal(t, "hello world", Add("hello ", "world"))
}

func TestCallList(t *testing.T) {
	Inc := func(i int) int {
		return i + 1
	}
	require.Equal(t, 4, CallList(2, Inc, Inc))
}

func TestMax(t *testing.T) {
	require.Equal(t, 2, Max(1, 2))
	require.Equal(t, 2, Max(2, 1))
}

func TestMin(t *testing.T) {
	require.Equal(t, 1, Min(1, 2))
	require.Equal(t, 1, Min(2, 1))
}
