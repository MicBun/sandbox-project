package web

import (
	"net/http"
	"sandbox/service"
	"sandbox/web/handlers"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		c.Set("session", auth)
		c.Next()
	}
}

func RegisterAPIRoutes(c *service.Container) {
	api := handlers.NewApiHandler(c)

	c.Web.GET("/hello", api.Hello)
	c.Web.POST("/login", api.Login)

	authRoutes := c.Web.Group("/")
	authRoutes.Use(Auth())
	authRoutes.GET("/orders", api.GetOrders)
	authRoutes.POST("/orders", api.PostOrder)
}
