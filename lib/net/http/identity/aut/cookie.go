package aut

import (
	"go-mall/lib/net/http/session"

	"github.com/gin-gonic/gin"
)

// AuthCookie AuthCookie model.
type AuthCookie struct {
	session session.Session
}

// NewCookieAuth new a cookie authentication.
func NewCookieAuth(c *gin.Context) *AuthCookie {
	return &AuthCookie{
		session: session.Default(c),
	}
}

// SignIn signin
// note: authentication must include the "NameIdentity" claim.
func (auth *AuthCookie) SignIn(claim Claim) {
	for k, v := range claim {
		if v != "" {
			auth.session.Set(string(k), v)
		}
	}
	auth.session.Save()
}

// SignOut signout
func (auth *AuthCookie) SignOut() {
	auth.session.Clear()
	auth.session.Save()
}

// IsAuthenticated whether has authenticated.
func (auth *AuthCookie) IsAuthenticated() bool {
	return auth.ClaimsIdentity() != ""
}

// ClaimsIdentity get identity claims. if not exists, return nil
func (auth *AuthCookie) ClaimsIdentity() string {
	return auth.Claim(NameIdentity)
}

// Claim get the claim value, if not exist, return zero value.
func (auth *AuthCookie) Claim(kind ClaimType) string {
	if v, ok := auth.session.Get(string(kind)).(string); ok {
		return v
	}
	return ""
}
