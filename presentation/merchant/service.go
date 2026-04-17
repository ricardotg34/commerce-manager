package merchant

import (
	"commerce-manager/config"
	"commerce-manager/domain/dtos"
	"commerce-manager/domain/entities"

	_ "gorm.io/driver/postgres"
)

func CreateMerchant(merchantDTO dtos.CreateMerchantDTO) error {
	var merchant = entities.Merchant{
		Name:      merchantDTO.Name,
		Commision: merchantDTO.Commision,
	}
	if err := config.DB.Create(&merchant).Error; err != nil {
		return err
	}
	return nil
}

func GetMerchantById(id uint) (entities.Merchant, error) {
	var merchant entities.Merchant
	result := config.DB.First(&merchant, id)
	return merchant, result.Error
}
