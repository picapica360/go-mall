package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Member 获取用户信息 API.
func (ctl Controller) Member(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, ctl.OK("member"))
}
