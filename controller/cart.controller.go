package controller

import "courseonline/services"

type CartController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewCartController(servicesManager services.ServiceManager) *CartController {
	return &CartController{
		serviceManager: &servicesManager,
	}
}
