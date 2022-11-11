package werr

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_withStack_Unwrap(t *testing.T) {
	fund := Errorf("%s", "test")
	require.Equal(t, "\"test\"", fmt.Sprintf("%q", fund))
	require.Equal(t, "test", fmt.Sprintf("%s", fund))

	err := WithMessagef(fund, "test error")
	stackErr := WithStack(err)
	require.NotEqual(t, stackErr, err)
	require.Equal(t, errors.Unwrap(stackErr), err)
	require.True(t, errors.Is(stackErr, err))
	require.Equal(t, fund, Cause(err))
	require.ErrorIs(t, stackErr, err)
	require.Equal(t, "test error: test", err.Error())
	require.Equal(t, "test error: test", fmt.Sprintf("%s", err))
	require.Greater(t, len(fmt.Sprintf("%+v", err)), len("test error: test"))
	require.Equal(t, "test error: test", fmt.Sprintf("%q", err))
}
