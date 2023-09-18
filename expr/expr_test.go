package expr

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

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
