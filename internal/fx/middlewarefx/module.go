package middlewarefx

import "go.uber.org/fx"

var Module = fx.Provide(
	provideMiddleware,
)
