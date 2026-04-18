package merchant

import (
	"commerce-manager/domain/dtos"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createMerchant(c *gin.Context) {
	logger := c.MustGet("logger").(*slog.Logger)
	var m dtos.CreateMerchantDTO
	if err := c.ShouldBindBodyWithJSON(&m); err != nil {
		dtos.Fail(c, http.StatusBadRequest, "E001", err.Error())
		return
	}
	res, err := CreateMerchant(m)
	if err != nil {
		dtos.Fail(c, http.StatusInternalServerError, "E002", "failed to create product")
	}

	logger.Info("create transaction",
		slog.Uint64("log_id", uint64(res.ID)),
		slog.String("action", "CREATE_MERCHANT"),
		slog.Float64("actor", 1),
		slog.Uint64("resource_id", uint64(res.ID)),
		slog.Time("timestamp", res.CreatedAt),
	)
	dtos.OK(c, gin.H{"message": "Commerce created"})
}

func getMerchantById(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		dtos.Fail(c, http.StatusBadRequest, "E003", "Merchant id not valid")
		return
	}
	merchant, err := GetMerchantById(uint(intId))
	if err != nil {
		dtos.Fail(c, http.StatusBadRequest, "E004", "Merchant id not found")
	}
	dtos.OK(c, merchant)
}

func updateMerchant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Commerce updated"})
}

func deleteMerchant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Commerce deleted"})
}
