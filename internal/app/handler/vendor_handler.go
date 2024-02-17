package handler

import (
	"net/http"
	"sf-delivery-delay-report/internal/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// VendorHandler struct
type VendorHandler struct {
	vendorService service.VendorService
}

// NewVendorHandler creates a new instance of VendorHandler
func NewVendorHandler(vendorService service.VendorService) *VendorHandler {
	return &VendorHandler{vendorService: vendorService}
}

// GetAllVendorsAPI API endpoint retrieves all vendors using the service and returns them as JSON
func (h *VendorHandler) GetAllVendorsAPI(c *gin.Context) {
	vendors, err := h.vendorService.GetAllVendors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vendors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vendors})
}

// GetOrdersByIDAPI API endpoint retrieves orders by vendor ID using the service and returns them as JSON
func (h *VendorHandler) GetOrdersByIDAPI(c *gin.Context) {
	vendorIDStr := c.Param("id")
	vendorID, err := strconv.ParseUint(vendorIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendor ID"})
		return
	}

	orders, err := h.vendorService.GetOrdersByID(uint(vendorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// GetDelayedOrdersByIDAPI API endpoint retrieves delayed orders by vendor ID using the service and returns them as JSON
func (h *VendorHandler) GetDelayedOrdersByIDAPI(c *gin.Context) {
	vendorIDStr := c.Param("id")
	vendorID, err := strconv.ParseUint(vendorIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendor ID"})
		return
	}

	orders, err := h.vendorService.GetDelayedOrdersByID(uint(vendorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve delayed orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}
