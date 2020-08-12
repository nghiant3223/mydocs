package middlewarefx

import (
	"github.com/nghiant3223/mydocs/pkg/middleware"
	"github.com/nghiant3223/mydocs/pkg/tokenmanager"
)

func provideMiddleware(manager tokenmanager.Manager) middleware.Middleware {
	return middleware.NewMiddleware(manager)
}
