package merchant

import "github.com/gin-gonic/gin"

func MerchantRoutes(merchant *gin.RouterGroup) {
	merchant.POST("/", createMerchant)
	merchant.GET("/:id", getMerchantById)
	merchant.PATCH("/:id", updateMerchant)
	merchant.DELETE("/:id", deleteMerchant)
}
