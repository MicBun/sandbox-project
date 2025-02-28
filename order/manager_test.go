package order

import (
	"sandbox/core"
	"sandbox/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSaveAndGetOrders(t *testing.T) {
	database.RunTest(func(db *gorm.DB) {
		om := NewManager(db)
		{
			orders, err := om.ListOrders(1, 0)
			assert.Error(t, err)
			assert.Len(t, orders, 0)
		}

		testOrder := core.Order{
			TrackingNumber:      "ABC123",
			ConsigneeAddress:    "Jalan Something",
			ConsigneeCity:       "Denpasar",
			ConsigneeProvince:   "Bali",
			ConsigneePostalCode: "1234",
			ConsigneeCountry:    "ID",
			Weight:              1,
			Height:              2,
			Width:               3,
			Length:              4,
		}
		err := om.SaveOrder(&testOrder)
		assert.NoError(t, err)

		testOrder2 := core.Order{
			TrackingNumber:      "ABC1234",
			ConsigneeAddress:    "Jalan Something",
			ConsigneeCity:       "Denpasar",
			ConsigneeProvince:   "Bali",
			ConsigneePostalCode: "1234",
			ConsigneeCountry:    "ID",
			Weight:              1,
			Height:              2,
			Width:               3,
			Length:              4,
		}
		err = om.SaveOrder(&testOrder2)
		assert.NoError(t, err)

		orders, err := om.ListOrders(1, 0)
		assert.NoError(t, err)
		assert.Len(t, orders, 1)

		orders2, err := om.ListOrders(2, 0)
		assert.NoError(t, err)
		assert.Len(t, orders2, 2)

		orders3, err := om.ListOrders(1, 1)
		assert.NoError(t, err)
		assert.Len(t, orders3, 1)

	})
}
