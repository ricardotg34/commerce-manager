package presentation

import (
	"commerce-manager/presentation/merchant"
	"commerce-manager/presentation/transaction"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	merchantRoutes := server.Group("/merchant")
	merchant.MerchantRoutes(merchantRoutes)

	transactionRoutes := server.Group("/transaction")
	transaction.TransactionRoutes(transactionRoutes)
}
