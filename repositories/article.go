package repositories

import (
	"github.com/nrad-K/blog-server/models"
	"gorm.io/gorm"
)

const (
	articleNumPerPage = 5
)

// 新規投稿をDBに登録する関数
func InsertArticle(db *gorm.DB, article models.Article) (models.Article, error) {
	newArticle := models.Article{
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		LikeNum:   article.LikeNum,
		CreatedAt: article.CreatedAt,
	}
	err := db.Create(&newArticle).Error
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

// 投稿一覧をDBから取得する関数
func SelectArticleList(db *gorm.DB, page int) ([]models.Article, error) {
	var articles []models.Article
	err := db.Limit(5).Offset((page - 1) * articleNumPerPage).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// 投稿IDを指定して、記事データを取得する関数
func SelectArticleDetail(db *gorm.DB, articleID int) (models.Article, error) {
	var article models.Article
	err := db.Where("article_id = ?", articleID).First(&article).Error
	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

// 投稿IDを指定して、記事データを削除する関数
func DeleteArticle(db *gorm.DB, articleID int) (models.Article, error) {
	var article models.Article
	tx := db.Begin()
	err := tx.Where("article_id = ?", articleID).First(&article).Error
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}
	err = tx.Where("article_id = ?", articleID).Delete(&models.Article{}).Error
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}
	tx.Commit()
	return article, nil
}

// 投稿内容をupdateする関数
func UpdateArticle(db *gorm.DB, updatedArticle models.Article) (models.Article, error) {
	var article models.Article
	tx := db.Begin()
	err := tx.Where("article_id = ? AND username = ?", updatedArticle.ID, updatedArticle.UserName).First(&article).Error
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}
	newArticle := models.Article{
		Title:    updatedArticle.Title,
		Contents: updatedArticle.Contents,
	}
	err = tx.Model(&article).Where("article_id = ? AND username = ?", updatedArticle.ID, updatedArticle.UserName).Updates(&newArticle).Error
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}
	tx.Commit()
	return article, nil
}

// いいねの数をupdateする関数
func UpdateLikeNum(db *gorm.DB, articleID int) (models.Article, error) {
	var article models.Article
	tx := db.Begin()
	err := tx.Where("article_id = ?", articleID).First(&article).Error
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}
	var niceNum int = article.LikeNum
	err = tx.Model(&article).Where("article_id = ?", articleID).Update("nice", niceNum+1).Error
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}
	tx.Commit()
	return article, nil
}
