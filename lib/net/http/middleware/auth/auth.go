package auth

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// CookieAuthConfig cookie config for auth.
type CookieAuthConfig struct {
	IgnoreRoutes       []string
	Domain             string
	Maxage             int
	Secure, HTTPOnly   bool
	AuthFailedRedirect string
}

// CookieAuth http cookie authentication middleware.
// ignoreRoutes -> ignore route colletion.
func CookieAuth(conf CookieAuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// test ignoreRoutes
		var hasIgnore bool
		if len(conf.IgnoreRoutes) > 0 {
			path := c.Request.URL.Path
			for _, route := range conf.IgnoreRoutes {
				if route != "" && checkIgnore(path, route) {
					hasIgnore = true
					break
				}
			}
		}

		if !hasIgnore {
			if _, err := c.Cookie(""); err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		c.Next()
	}
}

func checkIgnore(path, route string) bool {
	var (
		ok  bool
		err error
	)

	if ok, err = regexp.Match(route, []byte(path)); err != nil {
		return false
	}

	return ok
}
