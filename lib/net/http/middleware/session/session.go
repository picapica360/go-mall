package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Options session option.
type Options struct {
	Secret string
	Name   string

	Cookie *CookieOption
}

// CookieOption cookie options.
type CookieOption struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

// New a cookie session middleware.
func New(opt *Options) gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte(opt.Secret))
	if copt := opt.Cookie; copt != nil {
		sopt := sessions.Options{
			Path:     copt.Path,
			Domain:   copt.Domain,
			MaxAge:   copt.MaxAge,
			Secure:   copt.Secure,
			HttpOnly: copt.HTTPOnly,
		}
		store.Options(sopt)
	}

	return sessions.Sessions(opt.Name, store)
}
