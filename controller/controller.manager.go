package controller

import "courseonline/services"

type ControllerManager struct {
	*CourseController
	*UsersController
	*CategoryController
	*CourseimagesController
	*CourseodController
	*CartController
}

func NewControllerManager(store services.Store) *ControllerManager {
	return &ControllerManager{
		CourseController:       NewCourseController(store),
		UsersController:        NewUsersController(store),
		CategoryController:     NewCategoryController(store),
		CourseimagesController: NewCourseimagesController(store),
		CourseodController:     NewCourseodController(store),
		CartController:         NewCartController(store),
	}
}
