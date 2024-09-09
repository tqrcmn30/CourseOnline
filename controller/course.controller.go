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

// GetCourseById godoc
// @Summary 	GetcourseById
// @Description GetcourseById
// @Tags 		course
// @Accept  	json
// @Produce  	json
// @Param 		id    path 	int    		false 	"cours id"
// @Success 	200 	  {object} 	map[string]interface{}
// @Failure 	500 	  {} 	http.StatusInternalServerError
// @Failure 	404 	  {} 	http.StatusNotFound
// @Router /course/{id} [get]
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

// GetListCourse godoc
// @Summary      List course
// @Description  get course
// @Tags         course
// @Accept       json
// @Produce      json
// @Success      200  {object} 	map[string]interface{}
// @Failure      400  {} http.StatusInternalServerError
// @Failure      404  {} http.StatusInternalServerError
// @Failure      500  {} http.StatusInternalServerError
// @Security Bearer
// @Router       /course [get]
func (cour *CourseController) GetListCourse(c *gin.Context) {
	coursie, err := cour.storedb.GetAllCourses(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, coursie)
}

// PostCourse godoc
// @Summary		Create new course
// @Description	Create new course
// @Tags		course
// @Accept		json
// @Produce		json
// @Param		course	body		models.CoursePostReq	true	"course"
// @Success		201  {object} 	map[string]interface{}
// @Failure		422		{}	http.StatusUnprocessableEntity
// @Failure		500		{}	http.StatusInternalServerError
// @Security Bearer
// @Router	/course  [post]
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

// UpdateCourse godoc
// @Summary 	Update course
// @Description Update course
// @Tags 		course
// @Accept 		json
// @Produce 	json
// @Param 		id    path 	int    		false 	"cours id"
// @Param 		course 	body 		models.CourseUpdateReq 	true 	"course"
// @Success		201  {object} 	map[string]interface{}
// @Failure		422		{}	http.StatusUnprocessableEntity
// @Failure		500		{}	http.StatusInternalServerError
// @Failure 	404 	  {} 	http.StatusNotFound
// @Security 	Bearer
// @Router /course/{id} [put]
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

// DeleteCourse godoc
// @Summary Delete course
// @Description Delete course by id
// @Tags course
// @Accept json
// @Produce json
// @Param 		id    path 	int    		false 	"cours id"
// @Success 204
// @Failure		500		{}	http.StatusInternalServerError
// @Failure 	404 	{} 	http.StatusNotFound
// @Security Bearer
// @Router /course/{id} [delete]
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
