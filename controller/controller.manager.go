package controller

import "courseonline/services"

type ControllerManager struct {
	*CartController
	*CourseorderController
	*CourseController
	*UsersController
}

func NewControllerManager(serviceManager *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		CartController:        NewCartController(*serviceManager),
		CourseController:      NewCourseController(*serviceManager),
		CourseorderController: NewCourseorderController(*serviceManager),
		UsersController:       NewUsersController(*serviceManager),
	}
}
