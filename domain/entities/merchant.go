package entities

type Merchant struct {
	Model
	Name         string        `json:"name" example:"walmart"`
	Commision    float32       `json:"commision" example:"2.5"`
	Transactions []Transaction `json:"-"`
}
