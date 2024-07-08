package repository

import (
	"context"
	"database/sql"
	"errors"
	"pawpawchat/pkg/auth/model"
)

type AuthInfoRepository struct {
	db *sql.DB
}

func NewAuthInfoRepository(db *sql.DB) *AuthInfoRepository {
	return &AuthInfoRepository{db: db}
}

func (a *AuthInfoRepository) Create(ctx context.Context, authInfo *model.AuthInfo) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			err = errors.New("panic when executing a request")
		}

		if err != nil {
			tx.Rollback()
		}

		err = tx.Commit()
	}()

	query := `
		INSERT 
		INTO auth_info (email, hash_pass)
		VALUES ($1, $2)
		RETURNING user_id
	`
	authInfo.HashPass = authInfo.Password
	return a.db.QueryRowContext(ctx, query, authInfo.Email, authInfo.HashPass).Scan(&authInfo.UserID)
}
