package controller

import (
	"go-mall/app/admin/ums/service"

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
