package session

import (
	"go-mall/lib/net/http/session"

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
	store := session.NewCookieStore([]byte(opt.Secret))
	if copt := opt.Cookie; copt != nil {
		sopt := session.Options{
			Path:     copt.Path,
			Domain:   copt.Domain,
			MaxAge:   copt.MaxAge,
			Secure:   copt.Secure,
			HTTPOnly: copt.HTTPOnly,
		}
		store.Options(sopt)
	}

	return session.Sessions(opt.Name, store)
}
