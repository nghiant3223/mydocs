package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(g gin.IRouter)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{service: service}
}

func (c *controller) Register(g gin.IRouter) {
	g.GET("/items", c.getItemTree)
}

func (c *controller) getItemTree(ctx *gin.Context) {
	tree, err := c.service.GetItemTree()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, tree)
}
