package middleware

import (
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
	if len(tokens) == 0 {
		m.ReportError(ctx, apperrors.Unauthorized)
		return
	}

	token := tokens[0]
	prefix := token[:len(bearerPrefix)]
	content := token[len(bearerPrefix):]
	if prefix != bearerPrefix {
		m.ReportError(ctx, apperrors.InvalidToken)
		return
	}

	_, err := m.tokenManager.Validate(content)
	if err != nil {
		m.ReportError(ctx, apperrors.Unauthorized)
		return
	}
}
