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

func (o *OrderManager) ListOrders(filter core.OrderFilter) (orders []core.Order, err error) {
	q := o.db.Model(core.Order{})
	if filter.UserID != "" {
		q = q.Where("user_id = ?", filter.UserID)
	}
	err = q.Find(&orders).Error
	return orders, err
}
