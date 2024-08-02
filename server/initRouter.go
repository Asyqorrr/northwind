package server

import (
	"b30northwindapi/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(handler *controller.ControllerManager) *gin.Engine {
	router := gin.Default()

	// set a lower memory limit for multipart form
	router.MaxMultipartMemory = 8 << 20 //8Mib
	router.Static("/static", "./public")

	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin FB")
	})

	// GET = READ
	// POST = CREATE
	// PUT = UPDATE PARTIAL
	// PATCH = UPDATE 1 ATTRIBUTE
	// DELETE = DELETE

	categoryRoute := api.Group("/category")
	{
		categoryRoute.GET("", handler.GetListCategory)
		categoryRoute.GET("/", handler.GetListCategory)
		categoryRoute.GET("/:id", handler.GetCategoryById)
		categoryRoute.POST("/", handler.CreateCategory)
		categoryRoute.PUT("/:id", handler.UpdateCategory)
		categoryRoute.DELETE("/:id", handler.DeleteCategory)

	}

	productRoute := api.Group("/product")
	{
		productRoute.GET("", handler.FindAllProduct)
		productRoute.GET("/", handler.FindAllProduct)
		productRoute.POST("/", handler.CreateProduct)
		productRoute.GET("/paging", handler.FindAllProductPaging)
		productRoute.GET("/:id", handler.FindProductById)

		productRoute.POST("/multiUpload", handler.UploadMultipleProductImage)

		productRoute.PUT("/:id", handler.UpdateProduct)

		productRoute.DELETE("/:id", handler.DeleteProduct)
	}

	cartRoute := api.Group("/cart")
	{
		cartRoute.GET("/", handler.FindAllCart)
		cartRoute.GET("/:id", handler.FindCartById)
		cartRoute.GET("/paging", handler.FindAllCartPaging)
		
		cartRoute.POST("/", handler.CreateCart)

		cartRoute.PUT("/:id", handler.UpdateCart)
		
		cartRoute.DELETE("/:id", handler.DeleteCart)
	}

	return router
}
