package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Books only for test.
func (ctl Controller) Books(ctx *gin.Context) {
	books, _ := ctl.svc.Books(context.TODO())

	ctx.JSON(http.StatusOK, ctl.OK(books))
}

// Member 获取用户信息 API.
func (ctl Controller) Member(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctl.OK("member"))
}
