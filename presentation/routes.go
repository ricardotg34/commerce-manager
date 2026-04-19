package presentation

import (
	"commerce-manager/presentation/merchant"
	"commerce-manager/presentation/transaction"
	"commerce-manager/presentation/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	merchantRoutes := server.Group("/merchant")
	merchant.MerchantRoutes(merchantRoutes)

	transactionRoutes := server.Group("/transaction")
	transaction.TransactionRoutes(transactionRoutes)

	userRoutes := server.Group("/user")
	user.UserRoutes(userRoutes)
}
