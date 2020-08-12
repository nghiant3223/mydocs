package authfx

import (
	"github.com/nghiant3223/mydocs/internal/auth"
	"github.com/nghiant3223/mydocs/pkg/controller"
	"github.com/nghiant3223/mydocs/pkg/tokenmanager"
	"github.com/spf13/viper"
)

func provideController(manager tokenmanager.Manager) controller.Controller {
	password := viper.GetString("admin.password")

	config := auth.NewConfig(password)
	service := auth.NewService(config, manager)
	ctrl := auth.NewController(service)
	return ctrl
}
