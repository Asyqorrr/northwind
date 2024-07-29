package server

import (
	"b30northwindapi/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(handler *controller.ControllerManager) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin FB")
	})

	categoryRoute := api.Group("/category")
	{
		categoryRoute.GET("", handler.GetListCategory)
		categoryRoute.GET("/", handler.GetListCategory)
		categoryRoute.GET("/:id", handler.GetCategoryById)
		categoryRoute.POST("/", handler.CreateCategory)
		categoryRoute.PUT("/:id", handler.UpdateCategory)
		categoryRoute.DELETE("/:id", handler.DeleteCategory)

	}

	return router
}
