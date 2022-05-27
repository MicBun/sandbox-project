package web

import (
	"sandbox/service"
	"sandbox/web/handlers"
)

func RegisterAPIRoutes(c *service.Container) {
	api := handlers.NewApiHandler(c)

	c.Web.GET("/hello", api.Hello)
	c.Web.GET("/orders", api.GetOrders)
}
