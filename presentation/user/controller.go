package user

import (
	"commerce-manager/domain/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// postUser godoc
//
// @Summary		Create User
// @Tags	user
// @Description	This endpoint creates a new user
// @Accept		json
// @Produce		json
// @Param		body	body		dtos.CreateUserDTO		true	"Request body"
// @Success		200		{object}	dtos.SuccessResponse{data=entities.User}
// @Router			/user [post]
func postUser(c *gin.Context) {
	var u dtos.CreateUserDTO
	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		dtos.Fail(c, http.StatusBadRequest, "ERRU001", err.Error())
		return
	}
	res, err := CreateUser(u)
	if err != nil {
		dtos.Fail(c, http.StatusInternalServerError, "ERRU002", "failed to create product")
		return
	}
	dtos.OK(c, res)
}

// getUserById godoc
//
// @Summary		Get User
// @Tags	user
// @Description	This endpoint retreives a user by its id
// @Accept		json
// @Produce		json
// @Param		id	path	integer	true	"User id"
// @Success		200		{object}	dtos.SuccessResponse{data=entities.User}
// @Router			/user/{id} [get]
func getUserById(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		dtos.Fail(c, http.StatusBadRequest, "ERRU003", "User id not valid")
		return
	}
	user, err := GetUserByID(uint(intId))
	if err != nil {
		dtos.Fail(c, http.StatusBadRequest, "ERRU004", "User id not found")
		return
	}
	dtos.OK(c, user)
}
