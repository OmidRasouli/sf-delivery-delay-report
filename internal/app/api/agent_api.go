package api

import (
	"sf-delivery-delay-report/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterAgentAPI registers agent-related API endpoints
func RegisterAgentAPI(router *gin.Engine, agentHandler *handler.AgentHandler) {
	api := router.Group("/api/agent")

	// Assign delay report to agent API endpoint
	api.GET("/assign/:id", agentHandler.AssignDelayReportToAgent)

	// Mark delay report as done API endpoint
	api.GET("/done/:agent_id/:order_id", agentHandler.DelayReportDone)
}
