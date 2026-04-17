package dtos

type TransactionDTO struct {
	Amount     float32 `json:"amount" binding:"required,min=0"`
	UserID     uint    `json:"userID" binding:"required"`
	MerchantID uint    `json:"merchantID" binding:"required"`
}
