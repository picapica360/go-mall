package auth

import (
	"net/http"
	"regexp"

	"go-mall/lib/net/http/identity/aut"

	"github.com/gin-gonic/gin"
)

// CookieAuthOptions cookie options for auth.
type CookieAuthOptions struct {
	IgnoreRoutes []string
}

// CookieAuth http cookie authentication middleware.
// ignoreRoutes -> ignore route colletion.
func CookieAuth(opt CookieAuthOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var hasIgnore bool
		if len(opt.IgnoreRoutes) > 0 {
			path := c.Request.URL.Path
			for _, route := range opt.IgnoreRoutes {
				if route != "" && checkIgnore(path, route) {
					hasIgnore = true
					break
				}
			}
		}

		if !hasIgnore {
			auth := aut.NewCookieAuth(c)
			if !auth.IsAuthenticated() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errmsg": "authenticated failure."})
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

	if ok, err = regexp.MatchString(route, path); err != nil {
		return false
	}

	return ok
}
