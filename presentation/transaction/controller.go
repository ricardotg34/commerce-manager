package transaction

import (
	"commerce-manager/domain/dtos"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createTransaction(c *gin.Context) {
	logger := c.MustGet("logger").(*slog.Logger)

	var t dtos.TransactionDTO
	if err := c.ShouldBindBodyWithJSON(&t); err != nil {
		dtos.Fail(c, http.StatusBadRequest, "ERRT001", err.Error())
		return
	}
	res, err := ProcessTransaction(t)
	if err != nil {
		dtos.Fail(c, http.StatusInternalServerError, "ERRT002", err.Error())
	}
	logger.Info("create transaction",
		slog.Uint64("log_id", uint64(res.TransactionID)),
		slog.String("action", "CREATE_TRANSACTION"),
		slog.Float64("actor", float64(t.UserID)),
		slog.Uint64("resource_id", uint64(res.TransactionID)),
		slog.Time("timestamp", res.CreatedAt),
	)
	dtos.OK(c, res)
}
