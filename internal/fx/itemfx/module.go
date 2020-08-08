package itemfx

import "go.uber.org/fx"

var Module = fx.Provide(
	provideController,
)
