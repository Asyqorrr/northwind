package controller

import "b30northwindapi/services"

type ControllerManager struct {
	*categoryController
}

func NewControllerManager(serviceManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		categoryController: NewCategoryController(*serviceManager),
	}
}
