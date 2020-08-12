package dbfx

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nghiant3223/mydocs/internal/item"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func provideDB(lc fx.Lifecycle) (*gorm.DB, error) {
	dialect := viper.GetString("db.dialect")
	url := viper.GetString("db.url")
	log := viper.GetBool("db.log")

	db, err := gorm.Open(dialect, url)
	if err != nil {
		return nil, err
	}

	db.LogMode(log)
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return db.AutoMigrate(&item.Item{}).Error
			},
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
		},
	)
	return db, err
}
