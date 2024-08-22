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

type CourseorderController struct {
	storedb services.Store
}

func NewCourseorderController(store services.Store) *CourseorderController {
	return &CourseorderController{
		storedb: store,
	}
}

func (cuor *CourseorderController) GetListOrderCourse(c *gin.Context) {
	OrderCourse, err := cuor.storedb.GetAllOrderCourses(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, OrderCourse)
}

func (cuor *CourseorderController) PostOrderCourse(c *gin.Context) {
	var payload *models.OrderCoursePostReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := db.CreateOrderCourseParams{
		UscoPurchaseNo: payload.UscoPurchaseNo,
		UscoTax:        payload.UscoTax,
		UscoSubtotal:   payload.UscoSubtotal,
		UscoPatrxNo:    payload.UscoPatrxNo,
		UscoUserID:     payload.UscoUserID,
	}

	Courseorder, err := cuor.storedb.CreateOrderCourse(c, args)
	if err != nil {
		if apiErr := models.ConvertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(apiErr))
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, Courseorder)

}

func (cuor *CourseorderController) UpdateOrderCourse(c *gin.Context) {
	var payload *models.OrderCourseUpdateReq
	uscoId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateOrderCourseParams{
		UscoID:         int32(uscoId),
		UscoPurchaseNo: payload.UscoPurchaseNo,
		UscoTax:        payload.UscoTax,
		UscoSubtotal:   payload.UscoSubtotal,
		UscoPatrxNo:    payload.UscoPatrxNo,
		UscoUserID:     payload.UscoUserID,
	}

	Courseorder, err := models.Nullable(cuor.storedb.UpdateOrderCourse(c, *args))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusOK, Courseorder)

}

func (cuor *CourseorderController) DeleteOrderCourse(c *gin.Context) {
	uscoId, _ := strconv.Atoi(c.Param("id"))

	_, err := cuor.storedb.GetOrderCourseByID(c, int32(uscoId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = cuor.storedb.DeleteOrderCourse(c, int32(uscoId))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}
