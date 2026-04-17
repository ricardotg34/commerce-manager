package dtos

type CreateMerchantDTO struct {
	Name      string  `json:"name" binding:"required"`
	Commision float32 `json:"commision" binding:"required,min=0,max=100"`
}
