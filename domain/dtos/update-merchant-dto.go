package dtos

type UpdateMerchantDTO struct {
	Name      string  `json:"name,omitempty"`
	Commision float32 `json:"commision,omitempty" binding:",min=0,max=100"`
}
