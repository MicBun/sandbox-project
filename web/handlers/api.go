package handlers

import (
	"fmt"
	"net/http"
	"sandbox/core"
	"sandbox/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	ConsigneeName       string
	ConsigneeNumber     string
	ConsigneeCity       string
	ConsigneeProvince   string
	ConsigneePostalCode string
	ConsigneeCountry    string
	PaymentType         string
	Weight              float32
	Height              float32
	Width               float32
	Length              float32
}

func (h *apiHandler) GetOrders(c *gin.Context) {
	session, _ := c.Get("session")

	orders, err := h.container.OrdersManager.ListOrders(core.OrderFilter{
		UserID: session.(string),
	})
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	var resp []orderResponse
	for _, o := range orders {
		resp = append(resp, orderResponse{
			TrackingNumber:      o.TrackingNumber,
			ConsigneeAddress:    o.ConsigneeAddress,
			ConsigneeName:       o.ConsigneeName,
			ConsigneeNumber:     o.ConsigneeNumber,
			ConsigneeCity:       o.ConsigneeCity,
			ConsigneeProvince:   o.ConsigneeProvince,
			ConsigneePostalCode: o.ConsigneePostalCode,
			ConsigneeCountry:    o.ConsigneeCountry,
			PaymentType:         o.PaymentType,
			Weight:              o.Weight,
			Height:              o.Height,
			Width:               o.Width,
			Length:              o.Length,
		})
	}

	c.JSON(200, gin.H{"data": resp})
	return
}

type orderRequest struct {
	ConsigneeName       string `binding:"required"`
	ConsigneeNumber     string
	ConsigneeAddress    string  `json:"consigneeAddress" binding:"required"`
	ConsigneePostalCode string  `binding:"required"`
	ConsigneeCountry    string  `binding:"required"`
	ConsigneeCity       string  `binding:"required"`
	ConsigneeProvince   string  `binding:"required"`
	Length              float32 `binding:"required"`
	Width               float32 `binding:"required"`
	Height              float32 `binding:"required"`
	Weight              float32 `binding:"required"`
	PaymentType         string  `binding:"required,oneof='cod' 'prepaid'"`
}

func simpleResponse(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}

func (h *apiHandler) PostOrder(c *gin.Context) {
	var params orderRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation error",
				"details": simpleResponse(verr),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, _ := c.Get("session")

	newOrder := core.Order{
		TrackingNumber:      uuid.NewString(),
		ConsigneeName:       params.ConsigneeName,
		ConsigneeNumber:     params.ConsigneeNumber,
		ConsigneeAddress:    params.ConsigneeAddress,
		ConsigneePostalCode: params.ConsigneePostalCode,
		ConsigneeCity:       params.ConsigneeCity,
		ConsigneeProvince:   params.ConsigneeProvince,
		ConsigneeCountry:    params.ConsigneeCountry,
		PaymentType:         params.PaymentType,
		Weight:              params.Weight,
		Height:              params.Height,
		Width:               params.Width,
		Length:              params.Length,
		UserID:              session.(string),
	}

	err := h.container.OrdersManager.SaveOrder(&newOrder)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	resp := orderResponse{
		TrackingNumber:      newOrder.TrackingNumber,
		ConsigneeAddress:    newOrder.ConsigneeAddress,
		ConsigneeName:       newOrder.ConsigneeName,
		ConsigneeNumber:     newOrder.ConsigneeNumber,
		ConsigneeCity:       newOrder.ConsigneeCity,
		ConsigneeProvince:   newOrder.ConsigneeProvince,
		ConsigneePostalCode: newOrder.ConsigneePostalCode,
		ConsigneeCountry:    newOrder.ConsigneeCountry,
		PaymentType:         newOrder.PaymentType,
		Weight:              newOrder.Weight,
		Height:              newOrder.Height,
		Width:               newOrder.Width,
		Length:              newOrder.Length,
	}

	c.JSON(200, gin.H{"data": resp})
	return
}
