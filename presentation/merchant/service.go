package merchant

import (
	"commerce-manager/config"
	"commerce-manager/domain/dtos"
	"commerce-manager/domain/entities"

	_ "gorm.io/driver/postgres"
)

func CreateMerchant(merchantDTO dtos.CreateMerchantDTO) (*entities.Merchant, error) {
	var merchant = entities.Merchant{
		Name:      merchantDTO.Name,
		Commision: merchantDTO.Commision,
	}
	if err := config.DB.Create(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func GetMerchantById(id uint) (*entities.Merchant, error) {
	var merchant entities.Merchant
	result := config.DB.First(&merchant, id)
	return &merchant, result.Error
}

func UpdateMerchant(id uint, merchantDTO dtos.UpdateMerchantDTO) (*entities.Merchant, error) {
	merchant, err := GetMerchantById(id)
	if err != nil {
		return nil, err
	}

	config.DB.Model(merchant).Updates(entities.Merchant{
		Name:      merchantDTO.Name,
		Commision: merchantDTO.Commision,
	})

	return merchant, nil
}

func DeactivateMerchant(id uint) (*entities.Merchant, error) {
	merchant, err := GetMerchantById(id)
	if err != nil {
		return nil, err
	}

	config.DB.Delete(merchant)

	return merchant, nil
}
