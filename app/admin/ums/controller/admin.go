package controller

import (
	"net/http"

	"go-mall/app/admin/ums/model"

	"github.com/gin-gonic/gin"
)

// AdminRegister -> POST /register 管理员注册
func (ctl Controller) AdminRegister(c *gin.Context) {
	var input model.AdminParam
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, ctl.BadRequest(err))
		return
	}
	c.JSON(http.StatusOK, ctl.OK("member"))
}

// AdminLogin -> POST /login 管理员登录
func (ctl Controller) AdminLogin(c *gin.Context) {
	var input model.AdminParam
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, ctl.BadRequest(err))
		return
	}
	c.JSON(http.StatusOK, ctl.OK("member"))
}

// AdminLogout -> POST /logout 用户登出
func (ctl Controller) AdminLogout(c *gin.Context) {
	c.JSON(http.StatusOK, ctl.OKNull())
}

// AdminRole GET /role:id -> 获取管理者拥有的角色
func (ctl Controller) AdminRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ctl.BadRequest("必须传入参数 id"))
		return
	}
	c.JSON(http.StatusOK, ctl.OKNull())
}
