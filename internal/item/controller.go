package item

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
	"github.com/nghiant3223/mydocs/pkg/controller"
	"github.com/nghiant3223/mydocs/pkg/middleware"
	"github.com/spf13/cast"
)

type ctrl struct {
	controller.BaseController
	service    Service
	middleware middleware.Middleware
}

func NewController(service Service, middleware middleware.Middleware) controller.Controller {
	return &ctrl{service: service, middleware: middleware}
}

func (c *ctrl) Register(g gin.IRouter) {
	g.GET("/", c.getItemTree)
	g.GET("/:id", c.getOneItem)
	g.POST("/", c.middleware.VerifyToken, c.createItem)
	g.PATCH("/:id", c.middleware.VerifyToken, c.updateItem)
	g.DELETE("/:id", c.middleware.VerifyToken, c.deleteItem)
}

func (c *ctrl) getItemTree(ctx *gin.Context) {
	tree, err := c.service.GetItemTree()
	if err != nil {
		c.ReportError(ctx, err)
		return
	}
	c.ReportSuccess(ctx, tree)
}

func (c *ctrl) createItem(ctx *gin.Context) {
	var body CreateItemRequestBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		err = apperrors.InvalidItemData
		c.ReportError(ctx, err)
		return
	}
	item, err := c.service.CreateItem(body)
	if err != nil {
		c.ReportError(ctx, err)
		return
	}
	c.ReportSuccess(ctx, item)
}

func (c *ctrl) getOneItem(ctx *gin.Context) {
	paramID := ctx.Param("id")
	itemID := cast.ToUint(paramID)
	item, err := c.service.GetOneItem(itemID)
	if err != nil {
		c.ReportError(ctx, err)
		return
	}
	c.ReportSuccess(ctx, item)
}

func (c *ctrl) deleteItem(ctx *gin.Context) {
	paramID := ctx.Param("id")
	itemID := cast.ToUint(paramID)
	err := c.service.DeleteItem(itemID)
	if err != nil {
		c.ReportError(ctx, err)
		return
	}
	c.ReportSuccess(ctx, nil)
}

func (c *ctrl) updateItem(ctx *gin.Context) {
	paramID := ctx.Param("id")
	itemID := cast.ToUint(paramID)
	var body UpdateItemRequestBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		err = apperrors.InvalidItemData
		c.ReportError(ctx, err)
		return
	}
	item, err := c.service.UpdateItem(itemID, body)
	if err != nil {
		c.ReportError(ctx, err)
		return
	}
	c.ReportSuccess(ctx, item)
}
