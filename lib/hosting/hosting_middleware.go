package hosting

import (
	"go-mall/lib/net/http/middleware/session"
	"go-mall/lib/utils"

	"github.com/gin-gonic/gin"
)

// UseMiddleware add custom middleware to gin http pipe.
func (h *Host) UseMiddleware(middlewareFn ...func() gin.HandlerFunc) *Host {
	h.middlewareFn = append(h.middlewareFn, middlewareFn...)

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
