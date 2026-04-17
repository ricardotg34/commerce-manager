package presentation

import (
	"commerce-manager/presentation/merchant"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	merchantRoutes := server.Group("/merchant")
	merchant.MerchantRoutes(merchantRoutes)
}
