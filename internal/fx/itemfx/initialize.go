package itemfx

import (
	"github.com/jinzhu/gorm"
	"github.com/nghiant3223/mydocs/internal/item"
)

func provideController(db *gorm.DB) item.Controller {
	repo := item.NewRepository(db)
	service := item.NewService(repo)
	controller := item.NewController(service)
	return controller
}
