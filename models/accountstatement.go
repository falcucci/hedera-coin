package models

import "github.com/jinzhu/gorm"

// AccountStatement : AccountStatement of transactions
type AccountStatement struct {
	gorm.Model
	Type        int
	Description string
	Coin        int64
	ProductID   int
}
