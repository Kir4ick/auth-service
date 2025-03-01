package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"net/http"
)

var errorPostgresHttpCodeMap = map[string]int{
	"23505": http.StatusBadRequest,
	"23503": http.StatusConflict,
	"23502": http.StatusBadRequest,
	"22001": http.StatusBadRequest,
	"22007": http.StatusBadRequest,
	"22008": http.StatusBadRequest,
	"42883": http.StatusBadRequest,
	"42601": http.StatusInternalServerError,
	"40001": http.StatusServiceUnavailable,
	"08006": http.StatusServiceUnavailable,
}

type errorResponse struct {
	Message string `json:"message"`
}

func mapErrorToHttpStatusCode(err error) int {
	var pgErr *pgconn.PgError

	if errors.Is(err, pgx.ErrNoRows) {
		return http.StatusNotFound
	}

	if errors.As(err, &pgErr) {
		code, ok := errorPostgresHttpCodeMap[pgErr.Code]
		if ok {
			return code
		}
	}

	return http.StatusInternalServerError
}

func SendErrorResponse(code int, err error, ctx *gin.Context) {
	ctx.JSON(code, errorResponse{Message: err.Error()})
}

func SendErrorResponseByError(err error, ctx *gin.Context) {
	SendErrorResponse(mapErrorToHttpStatusCode(err), err, ctx)
}
