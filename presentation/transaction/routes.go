package transaction

import "github.com/gin-gonic/gin"

func TransactionRoutes(transaction *gin.RouterGroup) {
	transaction.POST("/", createTransaction)
}
