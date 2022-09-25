package account

import (
	"fmt"
	"net/http"

	"github.com/falcucci/maga-coin-api/config"
	magacoinpayment "github.com/falcucci/maga-coin-api/magacoin-payment"
	"github.com/falcucci/maga-coin-api/utils/response"
)

var (
	env = config.Env
)

type account struct {
	Number  string `json:"number"`
	Balance int64  `json:"balance"`
}

// GetBalance : Get account balance
func GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := env.TargetAccount
	balance, err := magacoinpayment.GetBalance(accountID)
	if err != nil {
		response.GenerateHTTPResponse(
			w, http.StatusInternalServerError,
			response.GenerateErrorResponse(response.InternalServerError,
				"Error", fmt.Sprintf("Failed to get account balance, err: %s", err)))
		return
	}
	response.GenerateHTTPResponse(
		w, http.StatusOK, response.GenerateSuccessResponse(balance.Records[0], 1, 1, 1))
}
