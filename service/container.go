package service

import (
	"sandbox/order"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Container struct {
	Web           *gin.Engine
	DB            *gorm.DB
	OrdersManager *order.OrderManager
}

func New(mainDB *gorm.DB) *Container {
	ginEngine := gin.Default()

	om := order.NewManager(mainDB)

	return &Container{
		Web:           ginEngine,
		DB:            mainDB,
		OrdersManager: om,
	}
}
