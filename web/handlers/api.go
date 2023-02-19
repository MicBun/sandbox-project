package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sandbox/service"
	"strconv"
)

type ApiHandlerInterface interface {
	Hello(c *gin.Context)
	GetOrders(c *gin.Context)
}

type apiHandler struct {
	container *service.Container
}

func NewApiHandler(container *service.Container) ApiHandlerInterface {
	return &apiHandler{
		container: container,
	}
}

// Hello godoc
// @Summary Hello
// @Description Hello
// @Tags Hello
// @Success 200 {object} map[string]interface{}
// @Router /hello [get]
func (h *apiHandler) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
	return
}

// orderResponse
type orderResponse struct {
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

// GetOrders godoc
//
//	@Summary		GetOrders
//	@Description	GetOrders
//	@Tags			GetOrders
//	@Produce		json
//	@Success		200	{object}	orderResponse
//	@Router			/orders [get]
func (h *apiHandler) GetOrders(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = -1
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = -1
	}
	orders, err := h.container.OrdersManager.ListOrders(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []orderResponse
	for _, o := range orders {
		resp = append(resp, orderResponse{
			ID:                  o.ID,
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

	c.JSON(200, gin.H{"data": resp, "count": len(resp)})
	return
}
