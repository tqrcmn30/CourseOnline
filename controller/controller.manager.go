package controller

import "courseonline/services"

type ControllerManager struct {
	*CartController
	*CourseorderController
	*CourseController
	*UsersController
	*CategoryController
	*CourseimagesController
	*CourseodController
}

func NewControllerManager(store services.Store) *ControllerManager {
	return &ControllerManager{
		CartController:         NewCartController(store),
		CourseController:       NewCourseController(store),
		CourseorderController:  NewCourseorderController(store),
		UsersController:        NewUsersController(store),
		CategoryController:     NewCategoryController(store),
		CourseimagesController: NewCourseimagesController(store),
		CourseodController:     NewCourseodController(store),
	}
}
