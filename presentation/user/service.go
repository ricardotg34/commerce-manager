package user

import (
	"commerce-manager/config"
	"commerce-manager/domain/entities"
)

func GetUserByID(id uint) (entities.User, error) {
	var user entities.User
	result := config.DB.First(&user, id)
	return user, result.Error
}
