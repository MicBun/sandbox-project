package order

import (
	"sandbox/core"

	"gorm.io/gorm"
)

type OrderManagerInterface interface {
	SaveOrder(order *core.Order) error
	ListOrders(limit int, offset int) ([]core.Order, error)
}

type OrderManager struct {
	Db *gorm.DB
}

func NewManager(db *gorm.DB) OrderManagerInterface {
	return &OrderManager{
		Db: db,
	}
}

func (o *OrderManager) SaveOrder(order *core.Order) error {
	return o.Db.Save(order).Error
}

func (o *OrderManager) ListOrders(limit int, offset int) (orders []core.Order, err error) {
	// err = o.Db.Model(core.Order{}).Find(&orders).Error
	err = o.Db.Model(core.Order{}).Limit(limit).Offset(offset).Find(&orders).Error

	return orders, err
}
