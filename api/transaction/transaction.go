package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/falcucci/maga-coin-api/api/accountstatement"
	"github.com/falcucci/maga-coin-api/config"
	magacoinpayment "github.com/falcucci/maga-coin-api/magacoin-payment"
	"github.com/falcucci/maga-coin-api/models"
	"github.com/falcucci/maga-coin-api/utils/response"
)

var (
	env = config.Env
)

// Request sasa
type Request struct {
	Amount    int64 `json:"amount"`
	Type      int   `json:"type"`
	ProductID int   `json:"product_id"`
}

// PostTransactions : Post of transactions
func PostTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		response.GenerateHTTPResponse(w, http.StatusInternalServerError, response.GenerateErrorResponse(response.InternalServerError,
			"Error", "Failed to start transaction transfer"))
		return
	}

	req := new(Request)
	requestBytes, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(requestBytes, req); err != nil {
		response.GenerateHTTPResponse(w, http.StatusInternalServerError, response.GenerateErrorResponse(response.InternalServerError,
			"Error", "Failed to build transaction transfer"))
		return
	}

	t := magacoinpayment.Transfer{
		PrivateKey:    env.PrivateKey,
		TargetAccount: env.TargetAccount,
		Amount:        req.Amount,
	}

	resp, err := magacoinpayment.PostTransfer(t, req.Type)
	if err != nil {
		response.GenerateHTTPResponse(w, http.StatusInternalServerError, response.GenerateErrorResponse(response.InternalServerError,
			"Error", fmt.Sprintf("Failed to perform transaction transfer, err: %s", err)))
		return
	}

	if resp.StatusCode == http.StatusPreconditionFailed {
		response.GenerateHTTPResponse(w, http.StatusPreconditionFailed,
			response.GenerateErrorResponse(response.InternalServerError,
				"Insufficient Payer Balance",
				"Was encountered an error when processing your request. We apologize for the inconvenience."))
		return
	}

	e := models.AccountStatement{
		Type:        req.Type,
		Description: "Transfer performed successfully!",
		Coin:        t.Amount,
		ProductID:   req.ProductID,
	}
	accountstatement.CreateAccountStatement(e)
	response.GenerateHTTPResponse(w, http.StatusOK, "Transfer performed successfully!")
	return
}
