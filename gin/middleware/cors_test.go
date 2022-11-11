package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewDefaultConfig(t *testing.T) {
	cfg := NewDefaultConfig(func(cfg *cors.Config) {
		cfg.MaxAge = 100
	})
	require.Equal(t, time.Duration(100), cfg.MaxAge)
}
