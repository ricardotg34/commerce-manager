package entities

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Transactions []Transaction
}
