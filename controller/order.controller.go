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

type OrderController struct {
	serviceManager *services.ServiceManager
}

func NewOrderController(servicesManager *services.ServiceManager) *OrderController {
	return &OrderController{
		serviceManager: servicesManager,
	}
}

type CreateOrderRequest struct {
	CustomerID     *string     `json:"customer_id" binding:"required"`
	EmployeeID     *int16      `json:"employee_id" binding:"required"`
	OrderDate      pgtype.Date `json:"order_date"`
	RequiredDate   pgtype.Date `json:"required_date"`
	ShippedDate    pgtype.Date `json:"shipped_date"`
	ShipVia        *int16      `json:"ship_via" binding:"required"`
	Freight        *float32    `json:"freight"`
	ShipName       *string     `json:"ship_name"`
	ShipAddress    *string     `json:"ship_address"`
	ShipCity       *string     `json:"ship_city"`
	ShipRegion     *string     `json:"ship_region"`
	ShipPostalCode *string     `json:"ship_postal_code"`
	ShipCountry    *string     `json:"ship_country"`
}

// func (handler *OrderController) DeleteCart(c *gin.Context) error {
// 	panic("not implemented") // TODO: Implement
// }

func (handler *OrderController) FindCartByCustomerAndProduct(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (handler *OrderController) FindCartByCustomerId(c *gin.Context) {
	id := c.Param("id")
	carts, err := handler.serviceManager.OrderService.FindCartByCustomerId(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	var response = &CartResponse{}
	response.CartId = int(carts[0].CartID)
	response.CustomerID = carts[0].CustomerID
	response.CompanyName = carts[0].CompanyName

	// fill carts data to dto response
	for _, v := range carts {
		product := &CartProductResponse{
			ProductID:   &v.ProductID,
			ProductName: &v.ProductName,
			UnitPrice:   v.UnitPrice,
			Qty:         v.Qty,
			Price:       v.Price,
		}
		response.Products = append(response.Products, product)
	}
	c.JSON(http.StatusOK, response)
}

func (handler *OrderController) FindCartByCustomerPaging(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (handler *OrderController) UpdateCartQty(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (handler *OrderController) FindOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	orders, err := handler.serviceManager.OrderService.FindOrderById(c, int16(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (handler *OrderController) FindAllOrder(c *gin.Context) {
	orders, err := handler.serviceManager.OrderService.FindAllOrder(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (handler *OrderController) CreateOrder(c *gin.Context) {
	var payload CreateOrderRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.CreateOrderParams{
		CustomerID: payload.CustomerID,
		EmployeeID: payload.EmployeeID,
		ShipVia:    payload.ShipVia,
	}

	newOrder, err := handler.serviceManager.OrderService.CreateOrderTx(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusCreated, newOrder)
}
