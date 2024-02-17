package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	var apiResponse APIResponse
	var err error

	response := makeRequest("POST", "/api/report-delay/1", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "order is not overdue yet", apiResponse.Msg)

	response = makeRequest("POST", "/api/report-delay/2", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "order is not overdue yet", apiResponse.Msg)

	response = makeRequest("POST", "/api/report-delay/3", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "order has been delivered", apiResponse.Msg)

	response = makeRequest("POST", "/api/report-delay/4", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "delay reported successfully", apiResponse.Msg)

	//Duplicate request for orderID 4
	response = makeRequest("POST", "/api/report-delay/4", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "delay reported successfully", apiResponse.Error)

	response = makeRequest("POST", "/api/report-delay/150", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusBadRequest, response.Code)

	response = makeRequest("POST", "/api/report-delay/5", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "order is not overdue yet", apiResponse.Msg)

	response = makeRequest("POST", "/api/report-delay/6", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "delay reported successfully", apiResponse.Error)

	response = makeRequest("POST", "/api/report-delay/7", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "order is not overdue yet", apiResponse.Msg)

	response = makeRequest("POST", "/api/report-delay/8", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "delay reported successfully", apiResponse.Error)

	response = makeRequest("POST", "/api/report-delay/9", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "order is not overdue yet", apiResponse.Msg)

	response = makeRequest("POST", "/api/report-delay/10", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "delay reported successfully", apiResponse.Error)

	response = makeRequest("POST", "/api/report-delay/11", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "delay reported successfully", apiResponse.Error)

	response = makeRequest("POST", "/api/report-delay/12", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "delay reported successfully", apiResponse.Error)

	response = makeRequest("POST", "/api/report-delay/13", nil)
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "order has been delivered", apiResponse.Msg)
}
