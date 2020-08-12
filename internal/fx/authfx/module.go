package authfx

import "go.uber.org/fx"

var Module = fx.Provide(
	fx.Annotated{
		Name:   "auth_controller",
		Target: provideController,
	},
)
