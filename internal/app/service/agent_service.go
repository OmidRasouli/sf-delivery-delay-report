package service

import (
	"sf-delivery-delay-report/internal/db/models"
	"sf-delivery-delay-report/internal/db/repository"
)

// AgentService interface defines the contract for agent-related operations
type AgentService interface {
	AssignDelayReportToAgent(agentID int) (*models.DelayReport, error)
	MarkDelayReportAsDone(agentID int, orderID uint) error
}

// AgentServiceImpl implements the AgentService interface
type AgentServiceImpl struct {
	agentRepository repository.AgentRepository
}

// NewAgentService creates a new instance of AgentService
func NewAgentService(agentRepository repository.AgentRepository) AgentService {
	return &AgentServiceImpl{agentRepository: agentRepository}
}

// AssignDelayReportToAgent assigns a delay report to an agent
func (s *AgentServiceImpl) AssignDelayReportToAgent(agentID int) (*models.DelayReport, error) {
	// Assign the delay report to the agent using the repository
	return s.agentRepository.AssignDelayReportToAgent(agentID)
}

// MarkDelayReportAsDone marks a delay report as done by the specified agent
func (s *AgentServiceImpl) MarkDelayReportAsDone(agentID int, orderID uint) error {
	// Mark the delay report as done using the repository
	return s.agentRepository.MarkDelayReportAsDone(agentID, orderID)
}
