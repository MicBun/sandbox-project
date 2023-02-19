package handlers_test

import (
	"encoding/json"
	"net/http"
	"sandbox/core"
	"sandbox/service"
	"sandbox/web"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		w, err := web.MakeRequest(c.Web, http.MethodGet, "/hello", nil)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			Message string
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Hello", resp.Message)
	})
}

func TestGetOrders(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		testOrder := func() []core.Order {
			var orders []core.Order
			for i := 0; i < 10; i++ {
				orders = append(orders, core.Order{
					TrackingNumber:      "ABC123" + fmt.Sprint(i),
					ConsigneeAddress:    "Jalan Something",
					ConsigneeCity:       "Denpasar",
					ConsigneeProvince:   "Bali",
					ConsigneePostalCode: "1234",
					ConsigneeCountry:    "ID",
					Weight:              1,
					Height:              2,
					Width:               3,
					Length:              4,
				})
			}
			return orders
		}()

		// testOrder := core.Order{
		// TrackingNumber:      "ABC123",
		// ConsigneeAddress:    "Jalan Something",
		// ConsigneeCity:       "Denpasar",
		// ConsigneeProvince:   "Bali",
		// ConsigneePostalCode: "1234",
		// ConsigneeCountry:    "ID",
		// Weight:              1,
		// Height:              2,
		// Width:               3,
		// Length:              4,
		// }
		for _, order := range testOrder {
			err := c.OrdersManager.SaveOrder(&order)
			assert.NoError(t, err)
		}
		// err := c.OrdersManager.SaveOrder(&testOrder[0])
		// assert.NoError(t, err)
		w, err := web.MakeRequest(c.Web, http.MethodGet, "/orders", nil)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			Data []struct {
				ID                  uint
				TrackingNumber      string
				ConsigneeAddress    string
				ConsigneeCity       string
				ConsigneeProvince   string
				ConsigneePostalCode string
				ConsigneeCountry    string
				Weight              float32
				Height              float32
				Width               float32
				Length              float32
			}
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Len(t, resp.Data, 1)
		assert.Equal(t, testOrder[0].TrackingNumber, resp.Data[0].TrackingNumber)
		assert.Equal(t, testOrder[0].ConsigneeAddress, resp.Data[0].ConsigneeAddress)
		assert.Equal(t, uint(1), resp.Data[0].ID)

		w, err = web.MakeRequest(c.Web, http.MethodGet, "/orders?limit=2&offset=0", nil)
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, testOrder[0].ConsigneeAddress, resp.Data[0].ConsigneeAddress)
		assert.Equal(t, uint(1), resp.Data[0].ID)
		assert.Equal(t, 2, len(resp.Data))

		w, err = web.MakeRequest(c.Web, http.MethodGet, "/orders?limit=1&offset=20", nil)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
