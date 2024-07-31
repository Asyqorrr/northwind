package controller

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewProductController(servicesManager services.ServiceManager) *ProductController{
	return &ProductController{
		serviceManager: &servicesManager,
	}
}

// Data transfer object
type CreateProductDto struct {
	ProductName     string   `form:"product_name" binding:"required"`
	SupplierID      *int16   `form:"supplier_id"`
	CategoryID      *int16   `form:"category_id"`
	QuantityPerUnit *string  `form:"quantity_per_unit"`
	UnitPrice       *float32 `form:"unit_price"`
	UnitsInStock    *int16   `form:"units_in_stock"`
	UnitsOnOrder    *int16   `form:"units_on_order"`
	ReorderLevel    *int16   `form:"reorder_level"`
	Discontinued    int32    `form:"discontinued"`
	ProductImage    *string  `form:"product_image"`
	Filename 		*SingleFileUpload 
}

type SingleFileUpload struct {
	Filename *multipart.FileHeader `form:"filename" binding:"required"`
}

func (handler *ProductController) CreateProduct(c *gin.Context){
	var payload *CreateProductDto

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


	args := &db.CreateProductParams{
		ProductName:     payload.ProductName,
		SupplierID:      payload.SupplierID,
		QuantityPerUnit: payload.QuantityPerUnit,
		UnitPrice:       payload.UnitPrice,
		UnitsInStock:    payload.UnitsInStock,
		UnitsOnOrder:    payload.UnitsOnOrder,
		Discontinued:    payload.Discontinued,
		ProductImage:    &newFileName,
	}

	product, err := handler.serviceManager.ProductService.CreateProduct(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (handler *ProductController) DeleteProduct(c *gin.Context)  {
	panic("not implemented") // TODO: Implement
}

func (handler *ProductController) FindAllProduct(c *gin.Context) {
	products, err := handler.serviceManager.ProductService.FindAllProduct(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	c.JSON(http.StatusOK, products)
}

func (handler *ProductController) FindAllProductPaging(c *gin.Context)  {
	// query from url
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	args := &db.FindAllProductPagingParams{
		Limit : int32(limit),
		Offset: int32(offset),
	}

	products, err := handler.serviceManager.ProductService.FindAllProductPaging(c, *args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	c.JSON(http.StatusOK, products)
}


func (handler *ProductController) FindProductById(c *gin.Context)  {
	panic("not implemented") // TODO: Implement
}


func (handler *ProductController) UpdateProduct(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
