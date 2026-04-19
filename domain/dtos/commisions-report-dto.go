package dtos

type MerchantCommisions struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type Report struct {
	Total     float64              `json:"total"`
	Merchants []MerchantCommisions `json:"merchant"`
}
