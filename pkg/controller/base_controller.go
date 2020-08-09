package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
)

const (
	defaultStatus = http.StatusInternalServerError
)

type BaseController struct{}

func (c *BaseController) HandleError(ctx *gin.Context, err error) {
	status := defaultStatus
	var e *apperrors.AppError
	if errors.As(err, &e) {
		status = e.Status
	}
	ctx.AbortWithError(status, e)
}
