package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
)

const (
	defaultStatus     = http.StatusOK
	returnCodeSuccess = 1
	returnCodeUnknown = 0
)

type BaseController struct{}

func (c *BaseController) ReportError(ctx *gin.Context, err error, message ...string) {
	var msg *string
	if len(message) > 0 {
		msg = &message[0]
	} else {
		errMsg := err.Error()
		msg = &errMsg
	}
	code := returnCodeUnknown
	var appError *apperrors.AppError
	if errors.As(err, &appError) {
		code = appError.Code
	}
	res := Response{
		ReturnCode: code,
		Message:    msg,
		Data:       nil,
	}
	ctx.JSON(defaultStatus, res)
}

func (c *BaseController) ReportSuccess(ctx *gin.Context, data interface{}, message ...string) {
	var msg *string
	if len(message) > 0 {
		msg = &message[0]
	}
	res := Response{
		ReturnCode: returnCodeSuccess,
		Message:    msg,
		Data:       data,
	}
	ctx.JSON(defaultStatus, res)
}
