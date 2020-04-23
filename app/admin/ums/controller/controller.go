package controller

import (
	"go-mall/app/admin/ums/service"
	"go-mall/lib/errcode"
	"go-mall/lib/mvc"
	"go-mall/lib/net/http/identity/aut"

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

// SignIn 用户登录。其中，若有多个角色，可使用逗号分隔； uid 不能为空值。
func (ctl *Controller) SignIn(c *gin.Context, uid, role string) {
	auth := aut.NewCookieAuth(c)
	claim := aut.Claim{
		aut.NameIdentity: uid,
		aut.Role:         role,
	}
	auth.SignIn(claim)
}

// SignOut 用户登出
func (ctl *Controller) SignOut(c *gin.Context) {
	auth := aut.NewCookieAuth(c)
	auth.SignOut()
}

// UID 用户唯一标识，如用户名、邮件地址或手机号等。
// 若没有值，则返回零值。
func (ctl *Controller) UID(c *gin.Context) string {
	auth := aut.NewCookieAuth(c)
	return auth.ClaimsIdentity()
}
