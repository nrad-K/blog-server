package services

import (
	"github.com/nrad-K/blog-server/errors"
	"github.com/nrad-K/blog-server/models"
	"github.com/nrad-K/blog-server/repositories"
)

// PostCommentHandlerで使用することを想定したサービス
// 引数の情報を元に新しい記事を作成後、結果を返却
func (s *AppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = errors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}
	return newComment, nil
}
