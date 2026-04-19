package dtos

type UpdateMerchantDTO struct {
	Name      string  `json:"name"`
	Commision float64 `json:"commision" binding:"min=0,max=100"`
}
