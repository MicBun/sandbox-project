package service

import (
	"sandbox/auth"
	"sandbox/order"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Container struct {
	Web           *gin.Engine
	DB            *gorm.DB
	OrdersManager order.OrderManagerInterface
	Auth          auth.AuthInterface
}

func New(mainDB *gorm.DB) *Container {
	ginEngine := gin.Default()

	om := order.NewManager(mainDB)
	auth := auth.CreateAuth(mainDB)

	return &Container{
		Web:           ginEngine,
		DB:            mainDB,
		OrdersManager: om,
		Auth:          auth,
	}
}
