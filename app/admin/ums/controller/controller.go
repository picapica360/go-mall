package controller

import (
	"go-mall/app/admin/ums/service"
	"go-mall/lib/errcode"
	"go-mall/lib/mvc"
	"go-mall/lib/net/http/authentication"

	"github.com/gin-gonic/gin"
)

// Config config
type Config struct {
	Svc service.Service
}

// Controller controller
type Controller struct {
	mvc.Controller

	svc service.Service
}

// New a controller
func New(c *Config) *Controller {
	return &Controller{
		svc: c.Svc,
	}
}

// BadRequest bad request， errcode 为 InputParamsError
func (ctl *Controller) BadRequest(err interface{}) map[string]interface{} {
	return ctl.BadCode(errcode.InputParamsError, err)
}

// SignIn 用户登录
func (ctl *Controller) SignIn(c *gin.Context, uid, role string) {
	auth := authentication.NewCookieAuth(c)
	claim := map[string]string{"uid": uid, "role": role}
	auth.SignIn(claim)
}

// SignOut 用户登出
func (ctl *Controller) SignOut(c *gin.Context) {
	auth := authentication.NewCookieAuth(c)
	auth.SignOut()
}

// UID 用户唯一标识，如用户名、邮件地址或手机号等。
// 若没有值，那么则返回零值。
func (ctl *Controller) UID(c *gin.Context) string {
	auth := authentication.NewCookieAuth(c)
	claim := auth.ClaimsIdentity()
	if claim != nil {
		if v, ok := claim["uid"]; ok {
			return v
		}
	}

	return ""
}
