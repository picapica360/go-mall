package controller

import (
	"go-mall/app/admin/ums/service"
	"go-mall/lib/errcode"

	"go-mall/lib/mvc"
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
