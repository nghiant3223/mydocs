package dbfx

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

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
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
			OnStart: func(ctx context.Context) error {
				return db.AutoMigrate(&item.Item{}).Error
			},
		},
	)
	return db, err
}
