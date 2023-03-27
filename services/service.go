package services

import (
	"gorm.io/gorm"
)

type AppService struct {
	db *gorm.DB
}

func NewAppService(db *gorm.DB) *AppService {
	return &AppService{db: db}
}
