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
		return
	}

	logger.Info("create merchant",
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
		return
	}
	dtos.OK(c, merchant)
}

func updateMerchant(c *gin.Context) {
	logger := c.MustGet("logger").(*slog.Logger)
	id := c.Param("id")
	intId, err := strconv.ParseUint(id, 10, 32)
	var m dtos.UpdateMerchantDTO
	if err := c.ShouldBindBodyWithJSON(&m); err != nil {
		dtos.Fail(c, http.StatusBadRequest, "E005", err.Error())
		return
	}
	res, err := UpdateMerchant(uint(intId), m)
	if err != nil {
		dtos.Fail(c, http.StatusInternalServerError, "E006", "failed to create product")
		return
	}
	logger.Info("update merchant",
		slog.Uint64("log_id", uint64(res.ID)),
		slog.String("action", "UPDATE_MERCHANT"),
		slog.Float64("actor", 1),
		slog.Uint64("resource_id", uint64(res.ID)),
		slog.Time("timestamp", res.UpdatedAt),
	)
	dtos.OK(c, res)
}

func deleteMerchant(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		dtos.Fail(c, http.StatusBadRequest, "E007", "Merchant id not valid")
		return
	}
	merchant, err := DeactivateMerchant(uint(intId))
	logger := c.MustGet("logger").(*slog.Logger)
	if err != nil {
		dtos.Fail(c, http.StatusBadRequest, "E008", "Merchant id not found")
		return
	}
	logger.Info("update merchant",
		slog.Uint64("log_id", intId),
		slog.String("action", "DELETE_MERCHANT"),
		slog.Float64("actor", 1),
		slog.Uint64("resource_id", uint64(merchant.ID)),
		slog.Time("timestamp", merchant.DeletedAt.Time),
	)
	dtos.OK(c, nil)
}
