package controller

import "courseonline/services"

type ControllerManager struct {
	*CartController
	*CourseorderController
	*CourseController
	*UsersController
}

func NewControllerManager(store services.Store) *ControllerManager {
	return &ControllerManager{
		CartController:        NewCartController(store),
		CourseController:      NewCourseController(store),
		CourseorderController: NewCourseorderController(store),
		UsersController:       NewUsersController(store),
	}
}
