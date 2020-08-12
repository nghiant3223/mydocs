package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
	"github.com/nghiant3223/mydocs/pkg/controller"
	"github.com/nghiant3223/mydocs/pkg/tokenmanager"
)

type Middleware interface {
	VerifyToken(c *gin.Context)
}

type middleware struct {
	controller.BaseController
	tokenManager tokenmanager.Manager
}

func NewMiddleware(manager tokenmanager.Manager) Middleware {
	return &middleware{tokenManager: manager}
}

func (m *middleware) VerifyToken(ctx *gin.Context) {
	tokens := ctx.Request.Header[authorization]
	if len(tokens) == 0 || !strings.HasPrefix(tokens[0], bearerPrefix) {
		m.ReportError(ctx, apperrors.Unauthorized)
		return
	}
	_, err := m.tokenManager.Validate(tokens[0])
	if err != nil {
		m.ReportError(ctx, apperrors.Unauthorized)
		return
	}
}
