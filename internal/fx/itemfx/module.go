package itemfx

import "go.uber.org/fx"

var Module = fx.Provide(
	fx.Annotated{
		Name:   "item_controller",
		Target: provideController,
	},
)
