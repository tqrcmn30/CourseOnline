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

type CourseodController struct {
	storedb services.Store
}

func NewCourseodController(store services.Store) *CourseodController {
	return &CourseodController{
		storedb: store,
	}
}

func (orde *CourseodController) GetListOrderCourse(c *gin.Context) {
	OrderCourse, err := orde.storedb.GetAllOrderCoursesDetails(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, OrderCourse)
}

func (orde *CourseodController) PostOrderCourseDetail(c *gin.Context) {
	var payload *models.OrderCoursesDetailPostReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := db.CreateOrderCoursesDetailParams{
		UcdeQty:        payload.UcdeQty,
		UcdeTotalPrice: payload.UcdeTotalPrice,
		UcdeUscoID:     payload.UcdeUscoID,
		UcdeCoursID:    payload.UcdeCoursID,
	}

	Courseod, err := orde.storedb.CreateOrderCoursesDetail(c, args)
	if err != nil {
		if apiErr := models.ConvertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(apiErr))
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, Courseod)

}

func (orde *CourseodController) UpdateOrderCourseDetail(c *gin.Context) {
	var payload *models.OrderCoursesDetailUpdateReq
	ucdeID, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateOrderCoursesDetailParams{
		UcdeID:         int32(ucdeID),
		UcdeQty:        payload.UcdeQty,
		UcdeTotalPrice: payload.UcdeTotalPrice,
		UcdeUscoID:     payload.UcdeUscoID,
		UcdeCoursID:    payload.UcdeCoursID,
	}

	Courseod, err := models.Nullable(orde.storedb.UpdateOrderCoursesDetail(c, *args))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusOK, Courseod)

}

func (orde *CourseodController) DeleteOrderCourseDetail(c *gin.Context) {
	ordeId, _ := strconv.Atoi(c.Param("id"))

	_, err := orde.storedb.GetOrderCourseByID(c, int32(ordeId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = orde.storedb.DeleteOrderCoursesDetail(c, int32(ordeId))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}
