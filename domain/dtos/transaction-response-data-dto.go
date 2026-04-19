package dtos

import "time"

type TransactionResponseDataDTO struct {
	TransactionID uint      `json:"transactionID" example:"1"`
	MerchantID    uint      `json:"merchantID" example:"1"`
	Amount        float64   `json:"amount" example:"1000"`
	Commision     float64   `json:"commision" example:"5.5"`
	Fee           float64   `json:"fee" example:"55"`
	CreatedAt     time.Time `json:"createdAt" example:"2026-01-01T00:00:00-06:00"`
}
