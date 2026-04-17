package transaction

import "github.com/gin-gonic/gin"

func CommerceRoutes(commerce *gin.RouterGroup) {
	commerce.POST("/", createTransaction)
}
