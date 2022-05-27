package handlers

import (
	"sandbox/service"

	"github.com/gin-gonic/gin"
)

type apiHandler struct {
	container *service.Container
}

func NewApiHandler(container *service.Container) *apiHandler {
	return &apiHandler{
		container: container,
	}
}

func (h *apiHandler) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
	return
}

type orderResponse struct {
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

func (h *apiHandler) GetOrders(c *gin.Context) {
	orders, err := h.container.OrdersManager.ListOrders()
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	var resp []orderResponse
	for _, o := range orders {
		resp = append(resp, orderResponse{
			TrackingNumber:      o.TrackingNumber,
			ConsigneeAddress:    o.ConsigneeAddress,
			ConsigneeCity:       o.ConsigneeCity,
			ConsigneeProvince:   o.ConsigneeProvince,
			ConsigneePostalCode: o.ConsigneePostalCode,
			ConsigneeCountry:    o.ConsigneeCountry,
			Weight:              o.Weight,
			Height:              o.Height,
			Width:               o.Width,
			Length:              o.Length,
		})
	}

	c.JSON(200, gin.H{"data": resp})
	return
}
