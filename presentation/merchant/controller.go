package merchant

import (
	"commerce-manager/domain/dtos"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// createMerchant godoc
//
// @Summary		Create Merchant
// @Tags	merchant
// @Description	This endpoint creates a new merchant
// @Accept		json
// @Produce		json
// @Param		body	body		dtos.CreateMerchantDTO		true	"Request body"
// @Success		200		{object}	dtos.SuccessResponse{data=entities.Merchant}
// @Router			/merchant [post]
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
		slog.String("log_id", uuid.New().String()),
		slog.String("action", "CREATE_MERCHANT"),
		slog.Float64("actor", 1),
		slog.Uint64("resource_id", uint64(res.ID)),
		slog.Time("timestamp", res.CreatedAt),
	)
	dtos.OK(c, res)
}

// getMerchantById godoc
//
// @Summary		Get Merchant
// @Tags	merchant
// @Description	This endpoint retreives a merchant by its id
// @Produce		json
// @Param		id	path	integer	true	"merchant id"
// @Success		200		{object}	dtos.SuccessResponse{data=entities.Merchant}
// @Router			/merchant/{id} [get]
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

// updateMerchant godoc
//
// @Summary		Update Merchant
// @Tags	merchant
// @Description	This endpoint updates either name and/or commision from a merchant
// @Accept		json
// @Produce		json
// @Param		body	body		dtos.UpdateMerchantDTO		true	"Request body"
// @Param		id	path	integer	true	"merchant id"
// @Success		200		{object}	dtos.SuccessResponse{data=entities.Merchant}
// @Router			/merchant/{id} [patch]
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
		slog.String("log_id", uuid.New().String()),
		slog.String("action", "UPDATE_MERCHANT"),
		slog.Float64("actor", 1),
		slog.Uint64("resource_id", uint64(res.ID)),
		slog.Time("timestamp", res.UpdatedAt),
	)
	dtos.OK(c, res)
}

// deleteMerchant godoc
//
// @Summary		Delete Merchant
// @Tags	merchant
// @Description	This endpoint deletes a merchant from the system (Soft delete)
// @Accept		json
// @Produce		json
// @Param		id	path	integer	true	"merchant id"
// @Success		200		{object}	dtos.SuccessResponse{data=nil}
// @Router			/merchant/{id} [delete]
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
		slog.String("log_id", uuid.New().String()),
		slog.String("action", "DELETE_MERCHANT"),
		slog.Float64("actor", 1),
		slog.Uint64("resource_id", uint64(merchant.ID)),
		slog.Time("timestamp", merchant.DeletedAt.Time),
	)
	dtos.OK(c, nil)
}
