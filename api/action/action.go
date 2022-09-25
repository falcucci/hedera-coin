package action

import (
	"net/http"

	"github.com/falcucci/maga-coin-api/models"
	"github.com/falcucci/maga-coin-api/utils/response"
)

var actions = []models.Action{
	models.Action{
		ID:          0,
		Description: "Venda Maga+",
		Coin:        100000000,
	},
	models.Action{
		ID:          1,
		Description: "Venda TV",
		Coin:        300000000,
	},
	models.Action{
		ID:          2,
		Description: "Venda cartão Luiza",
		Coin:        500000000,
	},
}

func mockActions() []models.Action {
	return actions
}

// GetDescriptionAction get description action by id
func GetDescriptionAction(ID int) string {
	for _, item := range actions {
		if item.ID == ID {
			return item.Description
		}
	}

	return "Sem descrição"
}

// GetActions : Return all actions
func GetActions(w http.ResponseWriter, r *http.Request) {
	actions := mockActions()
	response.GenerateHTTPResponse(w, http.StatusOK, response.GenerateSuccessResponse(actions, len(actions), 1, len(actions)))
}
