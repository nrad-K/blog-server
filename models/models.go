package models

import (
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id" gorm:"column:comment_id;type:integer;unsigned;autoIncrement;primaryKey"`
	ArticleID int       `json:"article_id" gorm:"column:article_id;type:integer;unsigned;not null;"`
	Message   string    `json:"message" gorm:"column:message;type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
}

type Article struct {
	ID          int       `json:"article_id" gorm:"column:article_id;type:integer;unsigned;autoIncrement;primaryKey;"`
	Title       string    `json:"title" gorm:"column:title;type:varchar(100) not null;"`
	Contents    string    `json:"contents" gorm:"column:contents;type:text;not null"`
	UserName    string    `json:"user_name" gorm:"column:username;type:varchar(100) not null;"`
	LikeNum     int       `json:"likes" gorm:"column:likes;type:integer;not null;"`
	CommentList []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;"`
}
