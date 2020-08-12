package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
	"github.com/nghiant3223/mydocs/pkg/controller"
)

type ctrl struct {
	controller.BaseController
	service Service
}

func NewController(service Service) controller.Controller {
	return &ctrl{service: service}
}

func (c *ctrl) Register(g gin.IRouter) {
	g.POST("/login", c.login)
}

func (c *ctrl) login(ctx *gin.Context) {
	var body LoginRequestBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		err = apperrors.InvalidLoginData
		c.ReportError(ctx, err)
		return
	}
	res, err := c.service.Login(body)
	if err != nil {
		c.ReportError(ctx, err)
		return
	}
	c.ReportSuccess(ctx, res)
}
