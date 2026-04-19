package dtos

type TransactionDTO struct {
	Amount     float64 `json:"amount" binding:"required,min=0" example:"1000"`
	UserID     uint    `json:"userID" binding:"required" example:"1"`
	MerchantID uint    `json:"merchantID" binding:"required" example:"1"`
}
