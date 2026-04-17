package transaction

import (
	"commerce-manager/domain/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createTransaction(c *gin.Context) {
	var t dtos.TransactionDTO
	if err := c.ShouldBindBodyWithJSON(&t); err != nil {
		dtos.Fail(c, http.StatusBadRequest, "ERRT001", err.Error())
		return
	}
	res, err := ProcessTransaction(t)
	if err != nil {
		dtos.Fail(c, http.StatusInternalServerError, "ERRT002", err.Error())
	}

	dtos.OK(c, res)
}
