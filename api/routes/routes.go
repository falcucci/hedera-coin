package routes

import (
	"github.com/falcucci/maga-coin-api/api/account"
	"github.com/falcucci/maga-coin-api/api/accountstatement"
	"github.com/falcucci/maga-coin-api/api/action"
	"github.com/falcucci/maga-coin-api/api/ping"
	"github.com/falcucci/maga-coin-api/api/reward"
	"github.com/falcucci/maga-coin-api/api/transaction"
	"github.com/gorilla/mux"
)

// Configure : Configure routes for API
func Configure() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", ping.Ping).Methods("GET")
	r.HandleFunc("/actions", action.GetActions).Methods("GET")
	r.HandleFunc("/account-statement", accountstatement.GetAccountStatement).Methods("GET")
	r.HandleFunc("/balance", account.GetBalance).Methods("GET")
	r.HandleFunc("/transactions", transaction.PostTransactions).Methods("POST")
	r.HandleFunc("/rewards", reward.GetRewards).Methods("GET")
	return r
}
