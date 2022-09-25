package magacoinpayment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/falcucci/maga-coin-api/config"
)

var (
	env = config.Env
)

// Transfer : aa
type Transfer struct {
	PrivateKey    string `json:"private_key"`
	TargetAccount int    `json:"target_account"`
	Amount        int64  `json:"amount"`
}

// Account : aa
type Account struct {
	Records []struct {
		Balance int64 `json:"balance"`
	} `json:"records"`
}

// PostTransfer - Post Magacoin transfer amount
func PostTransfer(transfer Transfer, transactionType int) (*http.Response, error) {
	payload, err := json.Marshal(transfer)
	if err != nil {
		return nil, err
	}

	fmt.Println("\n\npayload", string(payload))

	client := http.Client{
		Timeout: time.Second * time.Duration(env.MagacoinPaymentTimeout),
	}

	var url string

	if transactionType == 1 {
		url = fmt.Sprintf(
			"%s/wallet/cash-in",
			env.MagacoinPaymentURL)
	} else {
		url = fmt.Sprintf(
			"%s/wallet/cash-out",
			env.MagacoinPaymentURL)
	}

	fmt.Println("\n\nurl", url)

	req, err := http.NewRequest(
		"POST", url,
		bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		return resp, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf(
		"Unable to post payment transfer: %d - %s",
		resp.StatusCode, string(body))
}

// GetBalance - Get Magacoin account balance
func GetBalance(accountID int) (*Account, error) {
	client := http.Client{
		Timeout: time.Second * time.Duration(env.MagacoinPaymentTimeout),
	}

	url := fmt.Sprintf(
		"%s/wallet/balance/%d",
		env.MagacoinPaymentURL, accountID)

	req, err := http.NewRequest(
		"GET", url,
		bytes.NewBuffer(nil))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		balance := &Account{}
		err := json.Unmarshal(body, &balance)
		if err != nil {
			return nil, err
		}
		return balance, nil
	}

	return nil, fmt.Errorf(
		"Unable to get account balance: %d - %s",
		resp.StatusCode, string(body))
}
