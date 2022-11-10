package stringutil

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCharIsDigital(t *testing.T) {
	require.True(t, CharIsDigital('5'))
}
