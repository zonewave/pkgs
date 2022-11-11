package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrorString(t *testing.T) {
	err1 := FileNotFoundError{
		"test",
		"local",
	}
	require.Equal(t, "Config File \"test\" Not Found in \"local\"", err1.Error())
	err2 := UnsupportedConfigError("test")
	require.Equal(t, "Unsupported Config Type \"test\"", err2.Error())
	err3 := InvalidConfigTypeError("test")
	require.Equal(t, "test", err3.Error())
	err4 := ParseError{
		err3,
	}
	require.Equal(t, "While parsing config: test", err4.Error())

}
