package handler

import (
	"net/http"
	"sf-delivery-delay-report/internal/app/service"
	"sf-delivery-delay-report/internal/db/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AgentHandler struct
type AgentHandler struct {
	agentService service.AgentService
}

// NewAgentHandler creates a new instance of AgentHandler
func NewAgentHandler(agentService service.AgentService) *AgentHandler {
	return &AgentHandler{agentService: agentService}
}

// AssignDelayReportToAgent API endpoint assigns a delay report to an agent and returns the result as JSON
func (h *AgentHandler) AssignDelayReportToAgent(c *gin.Context) {
	agentIDStr := c.Param("id")
	agentID, err := strconv.ParseInt(agentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID"})
		return
	}

	reportDelay, err := h.agentService.AssignDelayReportToAgent(int(agentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": models.DelayReport{}, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reportDelay, "error": ""})
}

// DelayReportDone API endpoint marks a delay report as done for a specific agent and order, and returns the result as JSON
func (h *AgentHandler) DelayReportDone(c *gin.Context) {
	agentIDStr := c.Param("agent_id")
	orderIDStr := c.Param("order_id")
	agentID, err := strconv.ParseInt(agentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID"})
		return
	}
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.agentService.MarkDelayReportAsDone(int(agentID), uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark delay report as done"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "delay report marked as done"})
}
