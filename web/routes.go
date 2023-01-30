package web

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sandbox/service"
	"sandbox/web/handlers"
)

func RegisterAPIRoutes(c *service.Container) {
	api := handlers.NewApiHandler(c)

	c.Web.GET("/hello", api.Hello)
	c.Web.GET("/orders", api.GetOrders)

	// c.Web.GET("/pokemon", handlers.GetPokemon)

	c.Web.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
