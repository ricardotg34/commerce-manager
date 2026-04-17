package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Ammount    uint64
	Fee        float32
	UserID     uint
	MerchantID uint
}
