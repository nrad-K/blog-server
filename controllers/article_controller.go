package controllers

import (
	"github.com/nrad-K/blog-server/errors"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nrad-K/blog-server/controllers/services"
	"github.com/nrad-K/blog-server/models"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// POST /articles のハンドラ
func (c *ArticleController) PostArticleHandler(ctx *gin.Context) {
	var reqArticle models.Article
	if err := ctx.ShouldBindJSON(&reqArticle); err != nil {
		err = errors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		errors.ErrorHandler(ctx, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// PUT /articles のハンドラ
func (c *ArticleController) PutArticleHandler(ctx *gin.Context) {
	var reqArticle models.Article
	if err := ctx.ShouldBindJSON(&reqArticle); err != nil {
		err = errors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		errors.ErrorHandler(ctx, err)
		return
	}
	article, err := c.service.PutArticleService(reqArticle)
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// GET /articles/list のハンドラ
func (c *ArticleController) ArticleListHandler(ctx *gin.Context) {
	var page int
	var err error
	if page, err = strconv.Atoi(ctx.DefaultQuery("page", "1")); err != nil {
		err = errors.BadParam.Wrap(err, "query parameter must be a number")
		errors.ErrorHandler(ctx, err)
		return
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, articleList)
}

// GET /articles/{article_id} のハンドラ
func (c *ArticleController) ArticleDetailHandler(ctx *gin.Context) {
	var articleID int
	var err error
	if articleID, err = strconv.Atoi(ctx.Param("id")); err != nil {
		err = errors.BadParam.Wrap(err, "path parameter must be a number")
		errors.ErrorHandler(ctx, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, article)
}

// POST /articles/likes のハンドラ
func (c *ArticleController) PostLikeHandler(ctx *gin.Context) {
	var reqArticle models.Article
	if err := ctx.ShouldBindJSON(&reqArticle); err != nil {
		err = errors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		errors.ErrorHandler(ctx, err)
		return
	}

	article, err := c.service.PostLikeService(reqArticle)
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// DELETE /articles/{article_id} のハンドラ
func (c *ArticleController) DeleteArticleHandler(ctx *gin.Context) {
	var articleID int
	var err error
	if articleID, err = strconv.Atoi(ctx.Param("id")); err != nil {
		err = errors.BadParam.Wrap(err, "path parameter must be a number")
		errors.ErrorHandler(ctx, err)
		return
	}

	article, err := c.service.DeleteArticleService(articleID)
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"article_id": article.ID})
}
