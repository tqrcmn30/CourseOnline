package controller

import "courseonline/services"

type CourseodController struct {
	storedb services.Store
}

func NewCourseodController(store services.Store) *CourseodController {
	return &CourseodController{
		storedb: store,
	}
}
