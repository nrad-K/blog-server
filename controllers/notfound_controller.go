package controllers

import (
	"errors"

	apperrors "github.com/nrad-K/blog-server/errors"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler(ctx *gin.Context) {
	err := errors.New("no route")
	err = apperrors.ReqBodyDecodeFailed.Wrap(err, "not found")
	apperrors.ErrorHandler(ctx, err)
}
