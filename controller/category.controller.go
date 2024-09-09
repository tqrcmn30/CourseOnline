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

type CategoryController struct {
	storedb services.Store
}

func NewCategoryController(store services.Store) *CategoryController {
	return &CategoryController{
		storedb: store,
	}
}

// GetListCategory godoc
// @Summary      List category
// @Description  get category
// @Tags         category
// @Accept       json
// @Produce      json
// @Success      200  {object} 	map[string]interface{}
// @Failure      400  {} http.StatusInternalServerError
// @Failure      404  {} http.StatusInternalServerError
// @Failure      500  {} http.StatusInternalServerError
// @Security Bearer
// @Router       /category [get]
func (cate *CategoryController) GetListCategory(c *gin.Context) {
	categories, err := cate.storedb.GetAllCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, categories)
}

// PostCategory godoc
// @Summary		Create new category
// @Description	Create new category
// @Tags		category
// @Accept		json
// @Produce		json
// @Param		category	body		models.CategoryPostReq	true	"category"
// @Success		201  {object} 	map[string]interface{}
// @Failure		422		{}	http.StatusUnprocessableEntity
// @Failure		500		{}	http.StatusInternalServerError
// @Security Bearer
// @Router	/category  [post]
func (cate *CategoryController) PostCategory(c *gin.Context) {
	var payload *models.CategoryPostReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	category, err := cate.storedb.CreateCategory(c, payload.CateName)
	if err != nil {
		if apiErr := models.ConvertToApiErr(err); apiErr != nil {
			c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(apiErr))
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, category)

}

// UpdateCategory godoc
// @Summary 	Update category
// @Description Update category
// @Tags 		category
// @Accept 		json
// @Produce 	json
// @Param 		id    path 	int    		false 	"cate id"
// @Param 		category 	body 		models.CategoryUpdateReq 	true 	"category"
// @Success		201  {object} 	map[string]interface{}
// @Failure		422		{}	http.StatusUnprocessableEntity
// @Failure		500		{}	http.StatusInternalServerError
// @Failure 	404 	  {} 	http.StatusNotFound
// @Security 	Bearer
// @Router /category/{id} [put]
func (cate *CategoryController) UpdateCategory(c *gin.Context) {
	var payload *models.CategoryUpdateReq
	cateId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateCategoryParams{
		CateID:   int32(cateId),
		CateName: payload.CateName,
	}

	category, err := models.Nullable(cate.storedb.UpdateCategory(c, *args))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	if category == nil {
		c.JSON(http.StatusNotFound, models.NewError(models.ErrCategoryNotFound))
		return
	}
	c.JSON(http.StatusOK, category)

}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete category by id
// @Tags category
// @Accept json
// @Produce json
// @Param 		id    path 	int    		false 	"cate id"
// @Success 204
// @Failure		500		{}	http.StatusInternalServerError
// @Failure 	404 	{} 	http.StatusNotFound
// @Security Bearer
// @Router /category/{id} [delete]
func (cate *CategoryController) DeleteCategory(c *gin.Context) {
	cateId, _ := strconv.Atoi(c.Param("id"))

	_, err := cate.storedb.GetCategoryByID(c, int32(cateId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = cate.storedb.DeleteCategory(c, int32(cateId))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}

// GetCategoryById godoc
// @Summary 	GetCategoryById
// @Description GetCategoryById
// @Tags 		category
// @Accept  	json
// @Produce  	json
// @Param 		id    path 	int    		false 	"category id"
// @Success 	200 	  {object} 	map[string]interface{}
// @Failure 	500 	  {} 	http.StatusInternalServerError
// @Failure 	404 	  {} 	http.StatusNotFound
// @Router /category/{id} [get]
func (cate *CategoryController) GetCategoryById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := models.Nullable(cate.storedb.GetCategoryByID(c, int32(id)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	if category == nil {
		c.JSON(http.StatusNotFound, models.NewError(models.ErrCategoryNotFound))
		return
	}

	c.JSON(http.StatusOK, category)
}
