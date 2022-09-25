package dtos

import (
	"github.com/falcucci/maga-coin-api/api/action"
	"github.com/falcucci/maga-coin-api/api/reward"
	"github.com/falcucci/maga-coin-api/models"
)

// AccountStatement : AccountStatement DTO for response API
type AccountStatement struct {
	ID          int    `json:"id"`
	Type        int    `json:"type"`
	Description string `json:"description"`
	Coin        int64  `json:"coin"`
}

// MapModelToDto : Mapper model to DTO
func (e *AccountStatement) MapModelToDto(accountStatement models.AccountStatement) {
	e.ID = accountStatement.ID

	if accountStatement.Type == 1 {
		e.Description = action.GetDescriptionAction(accountStatement.ProductID)
	} else {
		e.Description = reward.GetDescriptionReward(accountStatement.ProductID)
	}

	e.Type = accountStatement.Type
	e.Coin = accountStatement.Coin
}
