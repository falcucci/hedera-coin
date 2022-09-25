package reward

import (
	"net/http"

	"github.com/falcucci/maga-coin-api/models"
	"github.com/falcucci/maga-coin-api/utils/response"
)

var rewards = []models.Reward{
	models.Reward{
		ID:            0,
		Description:   "Caneca Dev Monster",
		ExchangeValue: 500000000,
		Image:         "https://cdn.colab55.com/images/1485864048/studio/34095/art/70557/mugs/standard/standard/1.png",
	},
	models.Reward{
		ID:            1,
		Description:   "Visita ao Ensp e Labs São Paulo",
		ExchangeValue: 1000000000,
		Image:         "https://www.ymcafremont.org/wp-content/uploads/2018/12/Bus-trip.jpg",
	},
	models.Reward{
		ID:            2,
		Description:   "Skin vendedor supremo",
		ExchangeValue: 2000000000,
		Image:         "https://thumbs.dreamstime.com/z/avatar-de-um-guerreiro-do-jogo-v%C3%ADdeo-com-espada-96277763.jpg",
	},
}

func mockRewards() []models.Reward {
	return rewards
}

// GetRewards aa
func GetRewards(w http.ResponseWriter, r *http.Request) {
	rewards := mockRewards()
	response.GenerateHTTPResponse(w, http.StatusOK, response.GenerateSuccessResponse(rewards, len(rewards), 1, len(rewards)))
}

// GetDescriptionReward :
func GetDescriptionReward(ID int) string {
	for _, item := range rewards {
		if item.ID == ID {
			return item.Description
		}
	}

	return "Sem descrição"
}
