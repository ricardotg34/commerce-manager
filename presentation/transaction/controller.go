package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Transaction created"})
}
