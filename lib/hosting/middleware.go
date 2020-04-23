package hosting

import (
	"go-mall/lib/net/http/middleware/auth"
	"go-mall/lib/net/http/middleware/session"
	"go-mall/lib/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// UseMiddleware add custom middleware to gin http pipe.
// With the func arg, it will lazy execute only when be used.
func (h *Host) UseMiddleware(middlewareFn func() gin.HandlerFunc) *Host {
	h.middlewareFn = append(h.middlewareFn, middlewareFn)
	return h
}

// UseDefaultCookieSession add a cookie middleware by default config.
// the config from session node in app.[env].tool file.
func (h *Host) UseDefaultCookieSession() *Host {
	var fn = func() gin.HandlerFunc {
		cfg := h.conf.Session
		maxage, err := utils.AtoSecond(cfg.MaxAge)
		if err != nil {
			panic(err)
		}

		opt := session.Options{
			Secret: cfg.Secret,
			Name:   "picapica360",
			Cookie: &session.CookieOption{
				Path:     "/",
				Domain:   cfg.Domain,
				MaxAge:   maxage,
				Secure:   cfg.Secure,
				HTTPOnly: true,
			},
		}

		return session.New(&opt)
	}

	h.UseMiddleware(fn)
	return h
}

// UseCookieAuthentication use cookie authentication middleware to gin http pipe.
// ignoreRoutes -> route which not need be authenticated, regexp match.
func (h *Host) UseCookieAuthentication(ignoreRoutes ...string) *Host {
	var fn = func() gin.HandlerFunc {
		opt := auth.CookieAuthOptions{IgnoreRoutes: ignoreRoutes}
		return auth.CookieAuth(opt)
	}

	h.UseMiddleware(fn)
	return h
}

// UseCors use cors middleware to gin http pipe.
// origins -> the origins allowed, it will be “*” if is nil.
func (h *Host) UseCors(origins ...string) *Host {
	var fn = func() gin.HandlerFunc {
		config := cors.DefaultConfig()
		if len(origins) > 0 {
			config.AllowOrigins = origins
		} else {
			config.AllowAllOrigins = true
		}
		config.AllowCredentials = true // allow send cookie when cross site.
		return cors.New(config)
	}

	h.UseMiddleware(fn)
	return h
}

// UseCorsFn use cors middleware to gin http pipe.
func (h *Host) UseCorsFn(fn func(cors.Config) cors.Config) *Host {
	var mid = func() gin.HandlerFunc {
		config := fn(cors.DefaultConfig())
		return cors.New(config)
	}

	h.UseMiddleware(mid)
	return h
}
