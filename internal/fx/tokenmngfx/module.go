package tokenmngfx

import "go.uber.org/fx"

var Module = fx.Provide(
	provideTokenManager,
)
