package user

import "github.com/gin-gonic/gin"

func UserRoutes(user *gin.RouterGroup) {
	user.POST("/", postUser)
	user.GET("/:id", getUserById)
}
