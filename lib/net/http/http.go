package http

import (
	"go-mall/lib/config/env"
	"go-mall/lib/net/http/validator"

	"github.com/gin-gonic/gin"
)

// Config config
type Config struct {
	Handlers gin.HandlersChain // middlewares
}

// New a gin engine.
func New(c *Config) (engine *gin.Engine) {
	engine = makeEngine()
	engine.Use(builtinMiddleware()...) // add builtin middleware
	engine.Use(c.Handlers...)
	return
}

// Default create gin engine, with middleware.
func Default() (engine *gin.Engine) {
	engine = makeEngine()
	engine.Use(builtinMiddleware()...) // add builtin middleware
	return
}

func makeEngine() (engine *gin.Engine) {
	if env.IsProduction() {
		gin.SetMode("release") // debug or release
	}
	engine = gin.New()
	validator.Register() // register custom validators.
	return
}

func builtinMiddleware() []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	handlers = append(handlers, gin.Recovery())

	return handlers
}
