package repositories

import (
	"github.com/nrad-K/blog-server/models"
	"gorm.io/gorm"
)

// 新規投稿をinsertする関数
func InsertComment(db *gorm.DB, comment models.Comment) (models.Comment, error) {
	newComment := models.Comment{
		ArticleID: comment.ArticleID,
		Message:   comment.Message,
	}
	err := db.Create(&newComment).Error
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
func SelectCommentList(db *gorm.DB, articleID int) ([]models.Comment, error) {
	var comments []models.Comment
	err := db.Where("article_id = ?", articleID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
