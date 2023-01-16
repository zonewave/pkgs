package middleware

import (
	"github.com/zonewave/pkgs/util/cputil"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var defaultCORSConfig = cors.Config{
	AllowAllOrigins: true,
	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
	AllowHeaders: []string{
		"Origin",
		"Accept",
		"Accept-Language",
		"Content-Language",
		"Content-Type",
		"User-Agent",
		"x-user-id",
		"X-User-Id",
	},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}

// NewDefaultConfig  return a new Default Config object
func NewDefaultConfig(opts ...func(cfg *cors.Config)) *cors.Config {
	var cfg cors.Config
	_ = cputil.DeepCopy(&cfg, defaultCORSConfig)
	for _, opt := range opts {
		opt(&cfg)
	}
	return &cfg
}

// CORS enable CORS support
func CORS(configs ...*cors.Config) gin.HandlerFunc {
	if len(configs) != 0 {
		return cors.New(*configs[0])
	}

	return cors.New(defaultCORSConfig)
}
