package repository

import (
	"context"
	"pawpawchat/pkg/auth/model"

	"gorm.io/gorm"
)

type AuthInfoRepository struct {
	db *gorm.DB
}

func NewAuthInfoRepository(db *gorm.DB) *AuthInfoRepository {
	return &AuthInfoRepository{db: db}
}

func (a *AuthInfoRepository) Create(ctx context.Context, authInfo *model.AuthInfo) error {
	return a.db.Create(authInfo).Error
}
