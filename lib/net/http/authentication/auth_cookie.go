package authentication

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	DefaultAuthCookieName = "github.com/picapica360/go-mall"
)

type AuthCookie struct {
	context *gin.Context
}

// NewCookieAuth new a cookie authentication.
func NewCookieAuth(c *gin.Context) *AuthCookie {
	return &AuthCookie{
		context: c,
	}
}

// SignIn signin
func (a *AuthCookie) SignIn(claim ...Claim) {
	session := sessions.Default(a.context)
	session.Set(DefaultAuthCookieName, claim)
	session.Save()
}

// SignOut signout
func (a *AuthCookie) SignOut() {
	session := sessions.Default(a.context)
	session.Delete(DefaultAuthCookieName)
	session.Save()
}

// IsAuthenticated whether has authenticated.
func (a *AuthCookie) IsAuthenticated() bool {
	session := sessions.Default(a.context)
	v := session.Get(DefaultAuthCookieName)
	return v != nil
}
