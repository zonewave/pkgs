package di

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Option_String(t *testing.T) {

	opt := Provide(func() A {
		return A{"test"}
	})
	require.Equal(t, "fx.Provide(github.com/ybzhanghx/pkgs/di.Test_Option_String.func1())", opt.String())
	opt2 := Options(opt)
	require.Equal(t, "di.Options(fx.Provide(github.com/ybzhanghx/pkgs/di.Test_Option_String.func1()))", opt2.String())

	opt3 := validateOption{false}
	require.Equal(t, "fx.validate(false)", opt3.String())
}
