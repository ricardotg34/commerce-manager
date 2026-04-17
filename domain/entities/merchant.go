package entities

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name         string
	Commision    float32
	Transactions []Transaction
}
