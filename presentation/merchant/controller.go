package merchant

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createMerchant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Commerce created"})
}

func getMerchantById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Commerce retreived"})
}

func updateMerchant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Commerce updated"})
}

func deleteMerchant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Commerce deleted"})
}
