package main

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"github.com/nghiant3223/mydocs/internal/fx/dbfx"
	"github.com/nghiant3223/mydocs/internal/fx/itemfx"
	"github.com/nghiant3223/mydocs/internal/item"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	readConfig()
	app := fx.New(
		dbfx.Module,
		itemfx.Module,
		fx.Invoke(initialize),
	)
	app.Run()
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

func initialize(lc fx.Lifecycle, itemController item.Controller) {
	port := viper.GetString("port")

	router := gin.New()
	apiRouter := router.Group("/api")
	itemController.Register(apiRouter)

	err := router.Run(":" + port)
	if err != nil {
		log.Infof("Fail to start server on port " + port)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Infof("Server listening on port " + port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Infof("Server stopped")
			return nil
		},
	})
}
