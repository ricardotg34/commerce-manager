package transaction

import (
	"commerce-manager/config"
	"commerce-manager/domain/dtos"
	"commerce-manager/domain/entities"
	"commerce-manager/presentation/merchant"
	"commerce-manager/presentation/user"
	"errors"
	"math"

	_ "gorm.io/driver/postgres"
)

func ProcessTransaction(transactionDTO dtos.TransactionDTO) (*dtos.TransactionResponseDataDTO, error) {

	userO, err := user.GetUserByID(transactionDTO.UserID)
	if err != nil {
		return nil, errors.New("User not found")
	}

	merchantO, err := merchant.GetMerchantById(transactionDTO.MerchantID)
	if err != nil {
		return nil, errors.New("Merchant not found")
	}

	var fee = math.Round(transactionDTO.Amount * merchantO.Commision)

	var transaction = entities.Transaction{
		Ammount:    uint64(transactionDTO.Amount * 100),
		Fee:        uint64(fee),
		UserID:     userO.ID,
		MerchantID: merchantO.ID,
	}
	if err := config.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return &dtos.TransactionResponseDataDTO{
		TransactionID: transaction.ID,
		MerchantID:    transaction.MerchantID,
		Amount:        transactionDTO.Amount,
		Commision:     merchantO.Commision,
		Fee:           math.Round(fee) / 100,
		CreatedAt:     transaction.CreatedAt,
	}, nil
}

func RetreiveCommisionsReport() (*dtos.Report, error) {

	totalCommisions := entities.TotalCommisions{}
	merchantCommisions := []entities.MerchantCommisions{}

	if err := config.DB.Model(&entities.Transaction{}).Select("SUM(fee) AS total").Find(&totalCommisions).Error; err != nil {
		return nil, err
	}

	err := config.DB.Model(&entities.Transaction{}).
		Select("merchants.name, SUM(transactions.fee) AS total").
		Joins("RIGHT JOIN merchants ON transactions.merchant_id = merchants.id").
		Group("name").
		Find(&merchantCommisions).
		Error

	if err != nil {
		return nil, err
	}

	var merchantTotal []dtos.MerchantCommisions = make([]dtos.MerchantCommisions, len(merchantCommisions))

	for i, m := range merchantCommisions {
		merchantTotal[i] = dtos.MerchantCommisions{
			Name:  m.Name,
			Total: float64(m.Total) / 100,
		}
	}

	report := dtos.Report{
		Total:     float64(totalCommisions.Total) / 100,
		Merchants: merchantTotal,
	}

	return &report, nil

}
