package controller

import "courseonline/services"

type CourseController struct {
	storedb services.Store
}

func NewCourseController(store services.Store) *CourseController {
	return &CourseController{
		storedb: store,
	}
}
