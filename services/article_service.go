package services

import (
	apperrors "github.com/nrad-K/blog-server/errors"

	"errors"

	"github.com/nrad-K/blog-server/models"
	"github.com/nrad-K/blog-server/repositories"
	"gorm.io/gorm"
)

// ArticleDetailHandlerで使うことを想定したサービス
// 指定pageの記事詳細情報を返却
func (s *AppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostArticleHandlerで使うことを想定したサービス
// 引数の情報を元に新しい記事を作成後、返却
func (s *AppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

// PutArticleHandlerで使うことを想定したサービス
// 引数の情報を元に新しい記事を更新後、返却
func (s *AppService) PutArticleService(article models.Article) (models.Article, error) {
	updatedArticle, err := repositories.UpdateArticle(s.db, article)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update article")
		return models.Article{}, err
	}
	return models.Article{
		ID:        updatedArticle.ID,
		Title:     updatedArticle.Title,
		Contents:  updatedArticle.Contents,
		UserName:  updatedArticle.UserName,
		LikeNum:   updatedArticle.LikeNum,
		CreatedAt: updatedArticle.CreatedAt,
	}, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func (s *AppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}
	if len(articleList) == 0 {
		err = apperrors.NAData.Wrap(ErrNoData, "no data")
		return articleList, err
	}
	return articleList, nil
}

// PostLikeHandlerで使うことを想定したサービス
// 指定IDの記事のいいね数を+1して、結果を返却
func (s *AppService) PostLikeService(article models.Article) (models.Article, error) {
	updatedArticle, err := repositories.UpdateLikeNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}
	return models.Article{
		ID:        updatedArticle.ID,
		Title:     updatedArticle.Title,
		Contents:  updatedArticle.Contents,
		UserName:  updatedArticle.UserName,
		LikeNum:   updatedArticle.LikeNum,
		CreatedAt: updatedArticle.CreatedAt,
	}, nil
}

// DeleteArticleHandlerで使うことを想定したサービス
// 指定の投稿を削除して、結果を返却
func (s *AppService) DeleteArticleService(articleID int) (models.Article, error) {
	deletedArticle, err := repositories.DeleteArticle(s.db, articleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to delete target article")
		return models.Article{}, err
	}

	return models.Article{
		ID: deletedArticle.ID,
	}, nil
}
