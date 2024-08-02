package controller

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CartController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewCartController(servicesManager services.ServiceManager) *CartController{
	return &CartController{
		serviceManager: &servicesManager,
	}
}

// Data Transfer Object
type CreateCartsDto struct {
	CustomerID    string      `json:"customer_id"`
	ProductID     *int32      `json:"product_id"`
	UnitPrice     *float32    `json:"unit_price"`
	Qty           *int32      `json:"qty"`
	CartCreatedOn pgtype.Date `json:"cart_created_on"`
}

func (handler *CartController) CreateCart(c *gin.Context){
	var payload *CreateCartsDto

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.CreateCartsParams{
		CustomerID: 	payload.CustomerID,
		ProductID: 		payload.ProductID,
		UnitPrice: 		payload.UnitPrice,
		Qty: 			payload.Qty,
		CartCreatedOn: 	payload.CartCreatedOn,
	}

	cart, err := handler.serviceManager.CartService.CreateCarts(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusOK, cart)
}


func (handler *CartController) FindAllCart(c *gin.Context) {

	cart, err := handler.serviceManager.CartService.FindAllCarts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (handler *CartController) FindAllCartPaging(c *gin.Context)  {
	// query from url
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	args := &db.FindAllCartsPagingParams{
		Limit : int32(limit),
		Offset: int32(offset),
	}

	carts, err := handler.serviceManager.ProductService.FindAllCartsPaging(c, *args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	c.JSON(http.StatusOK, carts)
	
}

func (handler *CartController) FindCartById(c *gin.Context)  {
	cartId, _ := strconv.Atoi(c.Param("id"))

	cart, err := handler.serviceManager.CartService.FindCartsbyId(c, int32(cartId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (handler *CartController) DeleteCart(c *gin.Context)  {
	cartId, _ := strconv.Atoi(c.Param("id"))

	err := handler.serviceManager.CartService.DeleteCarts(c, int32(cartId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"data has been deleted"})
}

type UpdateCartDto struct {
	CartID        int32       `json:"cart_id"`
	CustomerID    string      `json:"customer_id"`
	ProductID     *int32      `json:"product_id"`
	UnitPrice     *float32    `json:"unit_price"`
	Qty           *int32      `json:"qty"`
	CartCreatedOn pgtype.Date `json:"cart_created_on"`
}


func (handler *CartController) UpdateCart(c *gin.Context) {
	var payload *UpdateCartDto
	
	cartId, _ := strconv.Atoi(c.Param("id"))
	
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewError(err))
		return
	}
	

	args := &db.UpdateCartsParams{
		CartID:      	  	int32(cartId),
		CustomerID:			payload.CustomerID,
		ProductID: 			payload.ProductID,
		UnitPrice: 			payload.UnitPrice,
		Qty: 				payload.Qty,
		CartCreatedOn: 		payload.CartCreatedOn,
	}

	cart, err := handler.serviceManager.CartService.UpdateCarts(c, *args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, cart)
}