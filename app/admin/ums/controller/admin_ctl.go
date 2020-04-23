package controller

import (
	"net/http"

	"go-mall/app/admin/ums/model"

	"github.com/gin-gonic/gin"
)

// AdminRegister -> POST /admin/register -> 管理员注册
func (ctl *Controller) AdminRegister(c *gin.Context) {
	var input model.AdminParam
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, ctl.BadRequest(err))
		return
	}
	c.JSON(http.StatusOK, ctl.OK("member"))
}

// AdminLogin -> POST /admin/login -> 管理员登录
func (ctl *Controller) AdminLogin(c *gin.Context) {
	type Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var input Login

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, ctl.BadRequest(err))
		return
	}

	// TODO: check login accout.
	ctl.SignIn(c, "picapica360", "admin,guest")

	c.JSON(http.StatusOK, ctl.OKNull())
}

// AdminLogout -> POST /admin/logout -> 用户登出
func (ctl *Controller) AdminLogout(c *gin.Context) {
	ctl.SignOut(c)

	c.JSON(http.StatusOK, ctl.OKNull())
}

// AdminInfo -> GET /admin/info -> 获取管理员信息
func (ctl *Controller) AdminInfo(c *gin.Context) {
	c.JSON(http.StatusOK, ctl.OKNull())
}

// AdminUpdate -> POST /admin/update/:id -> 管理人员信息更新
func (ctl *Controller) AdminUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, ctl.OKNull())
}

// AdminDelete -> POST /admin/delete/:id -> 管理人员信息删除
func (ctl *Controller) AdminDelete(c *gin.Context) {
	c.JSON(http.StatusOK, ctl.OKNull())
}

// AdminRole GET /admin/role/:id -> 获取管理者拥有的角色
func (ctl *Controller) AdminRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ctl.BadRequest("必须传入参数 id"))
		return
	}
	c.JSON(http.StatusOK, ctl.OKNull())
}
