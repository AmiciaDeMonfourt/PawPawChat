package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if email == "" {
		return nil, errors.New("missing email")
	}

	var user model.User
	query := `
		SELECT id, username
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user.Email = email

	return &user, nil
}

func (r *UserRepository) GetById(id uint64) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	query := `
		SELECT email, username  
		FROM users 
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Email, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user.ID = id

	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	query := `
		SELECT id, email 
		FROM users 
		WHERE username = $1
	`

	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user.Username = username

	return &user, nil
}
