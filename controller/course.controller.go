package controller

import (
	db "courseonline/db/sqlc"
	"courseonline/models"
	"courseonline/services"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	storedb services.Store
}

func NewCourseController(store services.Store) *CourseController {
	return &CourseController{
		storedb: store,
	}
}

func (cour *CourseController) GetCourseById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	course, err := models.Nullable(cour.storedb.GetCourseByID(c, int32(id)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	if course == nil {
		c.JSON(http.StatusNotFound, models.NewError(models.ErrCourseNotFound))
		return
	}

	c.JSON(http.StatusOK, course)
}

func (cour *CourseController) GetListCourse(c *gin.Context) {
	coursie, err := cour.storedb.GetAllCourses(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, coursie)
}

func (cour *CourseController) PostCourse(c *gin.Context) {
	var payload *models.CoursePostReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := db.CreateCourseParams{
		CoursName:   payload.CoursName,
		CoursDesc:   payload.CoursDesc,
		CoursAuthor: payload.CoursAuthor,
		CoursPrice:  payload.CoursPrice,
		CoursCateID: payload.CoursCateID,
	}

	Course, err := cour.storedb.CreateCourse(c, args)
	if err != nil {
		if apiErr := models.ConvertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(apiErr))
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, Course)

}

func (cour *CourseController) UpdateCourse(c *gin.Context) {
	var payload *models.CourseUpdateReq
	courId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateCourseParams{
		CoursID:     int32(courId),
		CoursName:   payload.CoursName,
		CoursDesc:   payload.CoursDesc,
		CoursAuthor: payload.CoursAuthor,
		CoursPrice:  payload.CoursPrice,
		CoursCateID: payload.CoursCateID,
	}

	Course, err := models.Nullable(cour.storedb.UpdateCourse(c, *args))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	if Course == nil {
		c.JSON(http.StatusNotFound, models.NewError(models.ErrCourseNotFound))
		return
	}

	c.JSON(http.StatusOK, Course)

}

func (cour *CourseController) DeleteCourse(c *gin.Context) {
	courId, _ := strconv.Atoi(c.Param("id"))

	_, err := cour.storedb.GetCourseByID(c, int32(courId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = cour.storedb.DeleteCourse(c, int32(courId))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}
