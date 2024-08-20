package controller

import "courseonline/services"

type UsersController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewUsersController(servicesManager services.ServiceManager) *UsersController {
	return &UsersController{
		serviceManager: &servicesManager,
	}
}
