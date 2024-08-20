package controller

import "courseonline/services"

type CourseorderController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewCourseorderController(servicesManager services.ServiceManager) *CourseorderController {
	return &CourseorderController{
		serviceManager: &servicesManager,
	}
}
