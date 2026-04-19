package entities

type User struct {
	ID           uint          `gorm:"primaryKey" json:"id" example:"1"`
	Name         string        `json:"name" example:"Ricardo"`
	Transactions []Transaction `json:"-"`
}
