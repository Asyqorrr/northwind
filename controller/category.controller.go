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

type CategoryController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewCategoryController(servicesManager services.ServiceManager) *CategoryController {
	return &CategoryController{
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







func (handler *CategoryController) UpdateCategory(c *gin.Context) {
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

	category, err := handler.serviceManager.CategoryService.UpdateCategory(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, category)

}








func (handler *CategoryController) DeleteCategory(c *gin.Context) {
	cateId, _ := strconv.Atoi(c.Param("id"))

	_, err := handler.serviceManager.CategoryService.FindCategoryById(c, int32(cateId))

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = handler.serviceManager.CategoryService.DeleteCategory(c, int32(cateId))
	if err != nil {
		
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})

}

func (handler *CategoryController) GetCategoryById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := handler.serviceManager.CategoryService.FindCategoryById(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, category)
}

// create category
func (handler *CategoryController) CreateCategory(c *gin.Context) {
	var payload *categoryCreateReq

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := db.CreateCategoryParams{
		CategoryName: payload.CategoryName,
		Description:  payload.Description,
	}

	category, err := handler.serviceManager.CategoryService.CreateCategory(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewValidationError(err))
		return
	}

	c.JSON(http.StatusCreated, category)
}

// get list category
func (handler *CategoryController) GetListCategory(c *gin.Context) {
	categories, err := handler.serviceManager.CategoryService.FindAllCategory(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrDataNotFound)
	}

	c.JSON(http.StatusOK, categories)
}
