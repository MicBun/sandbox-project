package order

import (
	"sandbox/core"

	"gorm.io/gorm"
)

type OrderManager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *OrderManager {
	return &OrderManager{
		db: db,
	}
}

func (o *OrderManager) SaveOrder(order *core.Order) error {
	return o.db.Save(order).Error
}

func (o *OrderManager) ListOrders() (orders []core.Order, err error) {
	err = o.db.Model(core.Order{}).Find(&orders).Error
	return orders, err
}
