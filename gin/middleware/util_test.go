package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_getHeader(t *testing.T) {
	c := &gin.Context{
		Request: &http.Request{Header: map[string][]string{"Accept": []string{"test"}}},
	}
	got, exist := getHeader(c, "Accept")
	require.True(t, exist)
	require.Equal(t, "test", got)
}
