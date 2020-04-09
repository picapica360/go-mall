package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Config session config
type Config struct {
	Secret string
	Name   string
}

// New a cookie session middleware.
func New(c *Config) gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte(c.Secret))
	return sessions.Sessions(c.Name, store)
}
