package itemfx

import (
	"github.com/jinzhu/gorm"
	"github.com/nghiant3223/mydocs/internal/item"
	"github.com/nghiant3223/mydocs/pkg/controller"
	"github.com/nghiant3223/mydocs/pkg/middleware"
)

func provideController(db *gorm.DB, middleware middleware.Middleware) controller.Controller {
	repo := item.NewRepository(db)
	service := item.NewService(repo)
	ctrl := item.NewController(service, middleware)
	return ctrl
}
