package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"pawpawchat/pkg/users/model"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
        INSERT INTO users (username, email, hash_pass) 
        VALUES ($1, $2, $3) 
        RETURNING id
    `

	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if email == "" {
		return nil, errors.New("missing email")
	}

	var user model.User
	query := `
		SELECT id, username, hash_pass
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.HashPass)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user.Email = email

	log.Printf("Users UserRepository: finding user: %v", user)

	return &user, nil
}
