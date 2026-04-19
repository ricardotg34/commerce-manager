package entities

type Transaction struct {
	Model
	Ammount    uint64
	Fee        uint64
	UserID     uint
	MerchantID uint
}
