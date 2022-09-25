package models

// Action : Action magalu
type Action struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Coin        int64  `json:"coin"`
}
