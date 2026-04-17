package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Ammount    int64
	Fee        int64
	UserID     uint
	MerchantID uint
}
