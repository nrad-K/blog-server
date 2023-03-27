package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nrad-K/blog-server/controllers/services"
	"github.com/nrad-K/blog-server/errors"
	"github.com/nrad-K/blog-server/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comments のハンドラ
func (c *CommentController) PostCommentHandler(ctx *gin.Context) {
	var reqComment models.Comment
	if err := ctx.ShouldBindJSON(&reqComment); err != nil {
		err = errors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		errors.ErrorHandler(ctx, err)
		return
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}
