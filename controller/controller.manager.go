package controller

import "b30northwindapi/services"

type ControllerManager struct {
	*CategoryController
	*ProductController
	*CartController
}

func NewControllerManager(serviceManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		CategoryController: NewCategoryController(*serviceManager),
		ProductController:  NewProductController(*serviceManager),
		CartController:  NewCartController(*serviceManager),
	}
}
