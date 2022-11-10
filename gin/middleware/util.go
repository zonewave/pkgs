package middleware

import (
	"github.com/gin-gonic/gin"
)

func getHeader(c *gin.Context, keys ...string) (string, bool) {
	for _, key := range keys {
		if v := c.GetHeader(key); v != "" {
			return v, true
		}
	}
	return "", false
}
