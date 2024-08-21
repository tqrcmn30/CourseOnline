package controller

import "courseonline/services"

type CourseorderController struct {
	storedb services.Store
}

func NewCourseorderController(store services.Store) *CourseorderController {
	return &CourseorderController{
		storedb: store,
	}
}
