package controller

import "courseonline/services"

type CourseController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewCourseController(servicesManager services.ServiceManager) *CourseController {
	return &CourseController{
		serviceManager: &servicesManager,
	}
}
