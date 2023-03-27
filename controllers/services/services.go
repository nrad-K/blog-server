package services

import "github.com/nrad-K/blog-server/models"

type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostLikeService(article models.Article) (models.Article, error)
	PutArticleService(article models.Article) (models.Article, error)
	DeleteArticleService(articleID int) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
