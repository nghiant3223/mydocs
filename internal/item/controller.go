package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
	"github.com/nghiant3223/mydocs/pkg/controller"
	"github.com/spf13/cast"
)

type itemController struct {
	controller.BaseController
	service Service
}

func NewController(service Service) controller.Controller {
	return &itemController{service: service}
}

func (c *itemController) Register(g gin.IRouter) {
	g.GET("/", c.getItemTree)
	g.GET("/:id", c.getOneItem)
	g.POST("/", c.createItem)
	g.PATCH("/:id", c.updateItem)
	g.DELETE("/:id", c.deleteItem)
}

func (c *itemController) getItemTree(ctx *gin.Context) {
	tree, err := c.service.GetItemTree()
	if err != nil {
		c.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, tree)
}

func (c *itemController) createItem(ctx *gin.Context) {
	var body CreateItemRequestBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		err = apperrors.InvalidItemData
		c.HandleError(ctx, err)
		return
	}
	item, err := c.service.CreateItem(body)
	if err != nil {
		c.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *itemController) getOneItem(ctx *gin.Context) {
	paramID := ctx.Param("id")
	itemID := cast.ToUint(paramID)
	item, err := c.service.GetOneItem(itemIgD)
	if err != nil {
		c.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *itemController) deleteItem(ctx *gin.Context) {
	paramID := ctx.Param("id")
	itemID := cast.ToUint(paramID)
	err := c.service.DeleteItem(itemID)
	if err != nil {
		c.HandleError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *itemController) updateItem(ctx *gin.Context) {
	paramID := ctx.Param("id")
	itemID := cast.ToUint(paramID)
	var body UpdateItemRequestBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		err = apperrors.InvalidItemData
		c.HandleError(ctx, err)
		return
	}
	item, err := c.service.UpdateItem(itemID, body)
	if err != nil {
		c.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}
