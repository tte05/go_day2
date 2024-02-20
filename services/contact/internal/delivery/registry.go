package delivery

import (
	"architecture_go/services/contact/internal/repository/controller"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

func (r registry) NewAppController() controller.AppController {
	return controller.AppController{
		Contact: r.NewContactController(),
	}
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}
