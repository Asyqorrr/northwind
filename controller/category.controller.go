package controller

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewCategoryController(servicesManager services.ServiceManager) *categoryController {
	return &categoryController{
		serviceManager: &servicesManager,
	}
}

// hold data from body client request
type categoryCreateReq struct {
	CategoryName string  `json:"category_name" binding:"required"`
	Description  *string `json:"description"`
}

type categoryUpdateReq struct {
	CategoryName string  `json:"category_name"`
	Description  *string `json:"description"`
}

func (handler *categoryController) UpdateCategory(c *gin.Context) {
	var payload *categoryUpdateReq
	cateId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.UpdateCategoryParams{
		CategoryID:   int32(cateId),
		CategoryName: payload.CategoryName,
		Description:  payload.Description,
	}

	category, err := handler.serviceManager.UpdateCategory(c, *args)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, category)

}

func (handler *categoryController) DeleteCategory(c *gin.Context) {
	cateId, _ := strconv.Atoi(c.Param("id"))

	_, err := handler.serviceManager.FindCategoryById(c, int32(cateId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = handler.serviceManager.DeleteCategory(c, int32(cateId))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}

func (handler *categoryController) GetCategoryById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := handler.serviceManager.FindCategoryById(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, category)
}

// create category
func (handler *categoryController) CreateCategory(c *gin.Context) {
	var payload *categoryCreateReq

	if err := c.ShouldBindJSON(&payload); err != nil {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "fail", "message": err})
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := db.CreateCategoryParams{
		CategoryName: payload.CategoryName,
		Description:  payload.Description,
	}

	category, err := handler.serviceManager.CreateCategory(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, category)

}

// get list category
func (handler *categoryController) GetListCategory(c *gin.Context) {
	categories, err := handler.serviceManager.FindAllCategory(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrDataNotFound)
	}

	c.JSON(http.StatusOK, categories)
}
