package authentication

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	// DefaultAuthCookieName default cookie name.
	DefaultAuthCookieName = "github.com/picapica360/go-mall"
)

// AuthCookie AuthCookie model.
type AuthCookie struct {
	session sessions.Session
}

// NewCookieAuth new a cookie authentication.
func NewCookieAuth(c *gin.Context) *AuthCookie {
	return &AuthCookie{
		session: sessions.Default(c),
	}
}

// SignIn signin
func (auth *AuthCookie) SignIn(claim Claim) {
	auth.session.Set(DefaultAuthCookieName, claim)
	auth.session.Save()
}

// SignOut signout
func (auth *AuthCookie) SignOut() {
	auth.session.Delete(DefaultAuthCookieName) // think: use Clear() to remove all cookie ?
	auth.session.Save()
}

// IsAuthenticated whether has authenticated.
func (auth *AuthCookie) IsAuthenticated() bool {
	v := auth.session.Get(DefaultAuthCookieName)
	return v != nil
}

// ClaimsIdentity get identity claims. if not exists, return nil
func (auth *AuthCookie) ClaimsIdentity() Claim {
	if v, ok := auth.session.Get(DefaultAuthCookieName).(Claim); ok {
		return v
	}
	return nil
}
