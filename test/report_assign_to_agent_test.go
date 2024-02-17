package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportAssignToAgent(t *testing.T) {
	var agentResponse AgentResponse
	var err error

	response := makeRequest("GET", "/api/agent/assign/1", nil)
	err = json.Unmarshal(response.Body.Bytes(), &agentResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, agentResponse.Data.AgentID)
	assert.Equal(t, uint(6), agentResponse.Data.OrderID)

	response = makeRequest("GET", "/api/agent/assign/1", nil)
	err = json.Unmarshal(response.Body.Bytes(), &agentResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "an unresolved report is already assigned to the agent", agentResponse.Error)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	response = makeRequest("GET", "/api/agent/assign/2", nil)
	err = json.Unmarshal(response.Body.Bytes(), &agentResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "record not found", agentResponse.Error)

	response = makeRequest("GET", "/api/agent/assign/200", nil)
	err = json.Unmarshal(response.Body.Bytes(), &agentResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "agent not found", agentResponse.Error)

	response = makeRequest("GET", "/api/agent/assign/3", nil)
	err = json.Unmarshal(response.Body.Bytes(), &agentResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "record not found", agentResponse.Error)

	response = makeRequest("GET", "/api/agent/assign/3", nil)
	err = json.Unmarshal(response.Body.Bytes(), &agentResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "record not found", agentResponse.Error)

	response = makeRequest("GET", "/api/agent/done/2/4", nil)
	var apiResponse APIResponse
	err = json.Unmarshal(response.Body.Bytes(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Failed to mark delay report as done", apiResponse.Error)

	response = makeRequest("GET", "/api/agent/assign/2", nil)
	err = json.Unmarshal(response.Body.Bytes(), &agentResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "record not found", agentResponse.Error)
}
