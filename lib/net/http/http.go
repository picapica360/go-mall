package http

import (
	"go-mall/lib/config/env"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

// Config config
type Config struct {
	Handlers gin.HandlersChain // middlewares
}

// New a gin engine.
func New(c *Config) (engine *gin.Engine) {
	handlers := builtinMiddleware()
	handlers = append(handlers, c.Handlers...)

	setGinMode()
	engine = gin.New()
	engine.Use(handlers...)
	return
}

// Default create gin engine, with middleware.
func Default() (engine *gin.Engine) {
	setGinMode()
	engine = gin.New()
	engine.Use(builtinMiddleware()...) // add middleware
	return
}

func setGinMode() {
	if env.IsProduction() {
		gin.SetMode("release") // debug or release
	}
}

func builtinMiddleware() []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, 0, 10)
	handlers = append(handlers, gin.Recovery(), cors.Default())
	// TODO: add gin middleware.
	return handlers
}
