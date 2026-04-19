package user

import (
	"commerce-manager/config"
	"commerce-manager/domain/dtos"
	"commerce-manager/domain/entities"
)

func CreateUser(userDTO dtos.CreateUserDTO) (*entities.User, error) {
	var user = entities.User{
		Name: userDTO.Name,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id uint) (entities.User, error) {
	var user entities.User
	result := config.DB.First(&user, id)
	return user, result.Error
}
