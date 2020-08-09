package itemfx

import (
	"github.com/jinzhu/gorm"
	"github.com/nghiant3223/mydocs/internal/item"
	"github.com/nghiant3223/mydocs/pkg/controller"
)

func provideController(db *gorm.DB) controller.Controller {
	repo := item.NewRepository(db)
	service := item.NewService(repo)
	ctrl := item.NewController(service)
	return ctrl
}
