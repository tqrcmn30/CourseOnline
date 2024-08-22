package controller

import "courseonline/services"

type CourseimagesController struct {
	storedb services.Store
}

func NewCourseimagesController(store services.Store) *CourseimagesController {
	return &CourseimagesController{
		storedb: store,
	}
}
