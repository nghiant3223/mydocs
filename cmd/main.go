package main

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"github.com/nghiant3223/mydocs/internal/fx/authfx"
	"github.com/nghiant3223/mydocs/internal/fx/dbfx"
	"github.com/nghiant3223/mydocs/internal/fx/itemfx"
	"github.com/nghiant3223/mydocs/internal/fx/middlewarefx"
	"github.com/nghiant3223/mydocs/internal/fx/tokenmngfx"
	"github.com/nghiant3223/mydocs/pkg/controller"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	readConfig()
	app := fx.New(
		dbfx.Module,
		itemfx.Module,
		authfx.Module,
		tokenmngfx.Module,
		middlewarefx.Module,
		fx.Invoke(initialize),
	)
	app.Run()
}

type InitParams struct {
	fx.In

	ItemController controller.Controller `name:"item_controller"`
	AuthController controller.Controller `name:"auth_controller"`
}

func initialize(lc fx.Lifecycle, params InitParams) {
	port := viper.GetString("port")

	router := gin.New()
	apiRouter := router.Group("/api")

	itemRouter := apiRouter.Group("/items")
	params.ItemController.Register(itemRouter)
	authRouter := apiRouter.Group("/auth")
	params.AuthController.Register(authRouter)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go startServer(router, port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Infof("Server stopped")
			return nil
		},
	})
}

func readConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("cannot read config"))
	}
}

func startServer(s *gin.Engine, port string) {
	err := s.Run(":" + port)
	if err != nil {
		log.Infof("Fail to start server on port " + port)
	}
}
