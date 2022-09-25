package models

// Reward djjd
type Reward struct {
	ID            int    `json:"id"`
	Description   string `json:"description"`
	ExchangeValue int    `json:"exchange_value"`
	Image         string `json:"image"`
}
