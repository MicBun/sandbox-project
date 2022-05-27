package handlers_test

import (
	"encoding/json"
	"net/http"
	"sandbox/core"
	"sandbox/service"
	"sandbox/web"
	"testing"

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
		err := c.OrdersManager.SaveOrder(&testOrder)
		assert.NoError(t, err)
		w, err := web.MakeRequest(c.Web, http.MethodGet, "/orders", nil)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			Data []struct {
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
		assert.Equal(t, testOrder.TrackingNumber, resp.Data[0].TrackingNumber)
		assert.Equal(t, testOrder.ConsigneeAddress, resp.Data[0].ConsigneeAddress)
	})
}
