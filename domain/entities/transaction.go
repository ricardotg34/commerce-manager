package entities

type Transaction struct {
	Model
	Ammount    uint64
	Fee        float32
	UserID     uint
	MerchantID uint
}
