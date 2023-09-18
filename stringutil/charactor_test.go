package stringutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCharIsDigital(t *testing.T) {
	require.True(t, CharIsDigital('5'))
}
