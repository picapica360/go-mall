package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-mall/lib/net/http/middleware/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCookieMiddleware(t *testing.T) {
	authMdw := auth.CookieAuthOptions{IgnoreRoutes: []string{"/test/"}}
	router := Default()
	router.Use(auth.CookieAuth(authMdw))
	router.GET("/test/cookie", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r := httptest.NewRequest(http.MethodGet, "/test/cookie", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, `response status code must be "200"`)
	assert.Equal(t, "OK", string(body), `response body must be "OK"`)
}
