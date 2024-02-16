package repository

import (
	"errors"
	"sf-delivery-delay-report/internal/db/models"

	"gorm.io/gorm"
)

// AgentRepository interface defines the contract for agent-related database operations
type AgentRepository interface {
	AssignDelayReportToAgent(agentID int) (*models.DelayReport, error)
	MarkDelayReportAsDone(agentID int, orderID uint) error
}

// AgentRepositoryImpl implements the AgentRepository interface
type AgentRepositoryImpl struct {
	db *gorm.DB
}

// NewAgentRepository creates a new instance of AgentRepository
func NewAgentRepository(db *gorm.DB) AgentRepository {
	return &AgentRepositoryImpl{db: db}
}

// AssignDelayReportToAgent assigns a delay report to an agent and returns the assigned delay report
func (r *AgentRepositoryImpl) AssignDelayReportToAgent(agentID int) (*models.DelayReport, error) {
	var agentExist models.Agent
	err := r.db.Where("id = ?", agentID).First(&agentExist).Error
	if err != nil {
		// Agent not found
		return nil, errors.New("agent not found")
	}

	var existingReport models.DelayReport
	err = r.db.Where("agent_id = ? AND solved = false", agentID).First(&existingReport).Error
	if err == nil {
		// An unresolved report is already assigned to the agent
		return nil, errors.New("an unresolved report is already assigned to the agent")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle other errors
		return nil, err
	}

	// Find the latest unsolved delay report
	var report models.DelayReport
	err = r.db.Where("solved = false and agent_id = -1").Order("created_at").First(&report).Error
	if err != nil {
		return nil, err
	}

	// Check if the delay report is already assigned
	if report.AgentID != -1 {
		// The delay report is already assigned
		return nil, errors.New("delay report is already assigned to an agent")
	}

	// Assign the delay report to the agent
	report.AgentID = agentID
	err = r.db.Save(&report).Error
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// MarkDelayReportAsDone marks a delay report as done by the specified agent
func (r *AgentRepositoryImpl) MarkDelayReportAsDone(agentID int, orderID uint) error {
	result := r.db.Model(&models.DelayReport{}).
		Where("agent_id = ? AND order_id = ? AND solved = false", agentID, orderID).
		Update("solved", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no matching records found for the given criteria")
	}

	return nil
}
