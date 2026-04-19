package entities

type Merchant struct {
	Model
	Name         string        `json:"name" example:"walmart"`
	Commision    float64       `json:"commision" example:"2.5"`
	Transactions []Transaction `json:"-"`
}
