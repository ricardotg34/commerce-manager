package dtos

import "time"

type TransactionResponseDataDTO struct {
	TransactionID uint      `json:"transactionID"`
	MerchantID    uint      `json:"merchantID"`
	Amount        float32   `json:"amount"`
	Commision     float32   `json:"commision"`
	Fee           float64   `json:"fee"`
	CreatedAt     time.Time `json:"createdAt"`
}
