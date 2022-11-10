package middleware

import (
	"github.com/gin-gonic/gin"
)

// Default returns with the Logger/Recovery/CORS/Locale/TimeOffset middleware already attached.
func Default() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
		CORS(),
	}
}
