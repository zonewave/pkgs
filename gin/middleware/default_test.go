package middleware

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDefault(t *testing.T) {
	g := gin.New()
	g.Use(Default()...)
	if len(g.Handlers) != 3 {
		t.Fatal("default middleware include 7 handlers")
	}
}
