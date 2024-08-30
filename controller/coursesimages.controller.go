package controller

import (
	db "courseonline/db/sqlc"
	"courseonline/models"
	"courseonline/services"
	"database/sql"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CourseimagesController struct {
	storedb services.Store
}

func NewCourseimagesController(store services.Store) *CourseimagesController {
	return &CourseimagesController{
		storedb: store,
	}
}

type CourseImagePostReq struct {
	//Filename    *SingleFileUpload
	Filename    []*multipart.FileHeader `form:"filename" binding:"required"`
	CoimDefault *string                 `form:"coim_default"`
	CoimRemeID  *int32                  `form:"coim_reme_id"`
}

type CourseImageUpdateReq struct {
	CoimID       int32   `form:"coim_id"`
	CoimFilename *string `form:"coim_filename"`
	Filename     *SingleFileUpload
	CoimDefault  *string `form:"coim_default"`
	CoimRemeID   *int32  `form:"coim_reme_id"`
}

type SingleFileUpload struct {
	Filename *multipart.FileHeader `form:"filename" binding:"required"`
}

type MultipleFileUpload struct {
	Filename []*multipart.FileHeader `form:"filename" binding:"required"`
}

func (ci *CourseimagesController) UploadMultipleProductImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	files := form.File["filename"]

	for _, v := range files {
		extension := filepath.Ext(v.Filename)
		// Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension

		// The file is received, so let's save it
		if err := c.SaveUploadedFile(v, "./public/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "message": "multiple product has been uploaded."})
}

func (ci *CourseimagesController) CreateCourseimages(c *gin.Context) {
	var payload *CourseImagePostReq

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	fileUpload, err := c.FormFile("filename")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(fileUpload.Filename)
	// Generate random file name for the new uploaded file
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(fileUpload, "./public/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	args := &db.CreateCourseImageParams{
		CoimFilename: &newFileName,
		CoimDefault:  payload.CoimDefault,
		CoimRemeID:   payload.CoimRemeID,
	}

	Courseimages, err := ci.storedb.CreateCourseImage(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, Courseimages)
}

func (ci *CourseimagesController) UpdateCourseImage(c *gin.Context) {
	var payload *CourseImageUpdateReq
	CoumId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateCourseImageParams{
		CoimID:       int32(CoumId),
		CoimFilename: payload.CoimFilename,
		CoimDefault:  payload.CoimDefault,
		CoimRemeID:   payload.CoimRemeID,
	}

	Courseimages, err := ci.storedb.UpdateCourseImage(c, *args)
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, Courseimages)
}

func (ci *CourseimagesController) FindCourseImageByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Courseimages, err := ci.storedb.GetCourseImageByID(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, Courseimages)
}

func (ci *CourseimagesController) DeleteCourseImage(c *gin.Context) {
	CoumId, _ := strconv.Atoi(c.Param("id"))

	_, err := ci.storedb.GetCourseImageByID(c, int32(CoumId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = ci.storedb.DeleteCourseImage(c, int32(CoumId))
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})
}

func (ci *CourseimagesController) GetListImage(c *gin.Context) {
	Courseimages, err := ci.storedb.GetAllCourseImages(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, Courseimages)
}
