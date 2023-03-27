package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}
	var statusCode int
	switch appErr.ErrCode {
	case NAData, NotFound:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	ctx.JSON(statusCode, appErr)
}
