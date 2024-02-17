package test

import (
	"net/http"
	"sf-delivery-delay-report/internal/db/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddOrder(t *testing.T) {
	const (
		orderCreateEndpoint = "/api/orders/create"
		tripCreateEndpoint  = "/api/trips/create"
	)
	//Add order 1
	order := models.Order{
		UserID:       1,
		DeliveryTime: time.Now().Add(time.Minute * 3),
		VendorID:     2,
	}
	response := makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 2
	order = models.Order{
		UserID:       4,
		DeliveryTime: time.Now().Add(time.Minute * 10),
		VendorID:     1,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 3
	order = models.Order{
		UserID:       6,
		DeliveryTime: time.Now().Add(time.Minute * -2),
		VendorID:     3,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 4
	order = models.Order{
		UserID:       2,
		DeliveryTime: time.Now().Add(time.Minute * -15),
		VendorID:     5,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 5
	order = models.Order{
		UserID:       8,
		DeliveryTime: time.Now().Add(time.Minute * 7),
		VendorID:     1,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 6
	order = models.Order{
		UserID:       7,
		DeliveryTime: time.Now().Add(time.Minute * -4),
		VendorID:     3,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 7
	order = models.Order{
		UserID:       9,
		DeliveryTime: time.Now().Add(time.Minute * 8),
		VendorID:     4,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 8
	order = models.Order{
		UserID:       5,
		DeliveryTime: time.Now().Add(time.Minute * -7),
		VendorID:     2,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 9
	order = models.Order{
		UserID:       10,
		DeliveryTime: time.Now().Add(time.Minute * 22),
		VendorID:     1,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 10
	order = models.Order{
		UserID:       9,
		DeliveryTime: time.Now().Add(time.Minute * -6),
		VendorID:     3,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 11
	order = models.Order{
		UserID:       2,
		DeliveryTime: time.Now().Add(time.Minute * 1),
		VendorID:     2,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 12
	order = models.Order{
		UserID:       7,
		DeliveryTime: time.Now().Add(time.Minute * 4),
		VendorID:     5,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 13
	order = models.Order{
		UserID:       3,
		DeliveryTime: time.Now().Add(time.Minute * -3),
		VendorID:     1,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 14
	order = models.Order{
		UserID:       1,
		DeliveryTime: time.Now().Add(time.Minute * -2),
		VendorID:     5,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 15
	order = models.Order{
		UserID:       4,
		DeliveryTime: time.Now().Add(time.Minute * -3),
		VendorID:     5,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 16
	order = models.Order{
		UserID:       9,
		DeliveryTime: time.Now().Add(time.Minute * 10),
		VendorID:     2,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 17
	order = models.Order{
		UserID:       6,
		DeliveryTime: time.Now().Add(time.Minute * 17),
		VendorID:     4,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 18
	order = models.Order{
		UserID:       10,
		DeliveryTime: time.Now().Add(time.Minute * -9),
		VendorID:     1,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 19
	order = models.Order{
		UserID:       3,
		DeliveryTime: time.Now().Add(time.Minute * 1),
		VendorID:     5,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add order 20
	order = models.Order{
		UserID:       7,
		DeliveryTime: time.Now().Add(time.Minute * 12),
		VendorID:     4,
	}
	response = makeRequest("POST", orderCreateEndpoint, order)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add trip for order 3
	trip := models.Trip{
		OrderID: 3,
		Status:  "DELIVERED",
	}
	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add trip for order 4
	trip = models.Trip{
		OrderID: 4,
		Status:  "PICKED",
	}
	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)

	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add trip for order 8
	trip = models.Trip{
		OrderID: 8,
		Status:  "ASSIGNED",
	}
	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add trip for order 10
	trip = models.Trip{
		OrderID: 10,
		Status:  "ASSIGNED",
	}
	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add trip for order 13
	trip = models.Trip{
		OrderID: 13,
		Status:  "DELIVERED",
	}
	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add trip for order 14
	trip = models.Trip{
		OrderID: 14,
		Status:  "PICKED",
	}
	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)

	//Add trip for order 18
	trip = models.Trip{
		OrderID: 18,
		Status:  "ASSIGNED",
	}
	response = makeRequest("POST", tripCreateEndpoint, trip)
	assert.Equal(t, http.StatusOK, response.Code)
}
