package controller

import "b30northwindapi/services"

type ControllerManager struct {
	*CategoryController
	*ProductController
	*CartController
	*OrderController
}

func NewControllerManager(serviceManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		CategoryController: NewCategoryController(*serviceManager),
		ProductController:  NewProductController(*serviceManager),
		CartController:  NewCartController(*serviceManager),
		OrderController:  NewOrderController(serviceManager),
	}
}
