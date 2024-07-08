package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	if user.FirstName == "" || user.SecondName == "" {
		return fmt.Errorf("missing user data, requests user: %v", user)
	}

	tx, err := r.db.BeginTx(ctx, nil)
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

	insertIntoUsers := `
		INSERT INTO users (id, username, first_name, second_name)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`
	return r.db.QueryRowContext(ctx, insertIntoUsers, user.ID, user.Username, user.FirstName, user.SecondName).Scan(&user.ID)
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if email == "" {
	// 	return nil, errors.New("missing email")
	// }

	// var user model.User
	// query := `
	// 	SELECT id, username, first_name, second_name
	// 	FROM users
	// 	WHERE email = $1
	// `

	// err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.FirstName, &user.SecondName)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, nil
	// 	}
	// 	return nil, err
	// }

	// user.Email = email

	return nil, nil
}

func (r *UserRepository) GetById(id uint64) (*model.User, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// var user model.User

	// getUserQuery := `
	// 	WITH auth AS(
	// 		SELECT user_id, email, last_login
	// 	)
	// 	SELECT users.*, auth.*
	// 	FROM users u
	// 	JOIN auth ON users.id = auth.user_id
	// `

	// res := r.db.QueryRowContext(ctx, getUserQuery)
	// res.Scan()
	return nil, nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// var user model.User
	// query := `
	// 	SELECT id, email
	// 	FROM users
	// 	WHERE username = $1
	// `

	// err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Email)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, nil
	// 	}
	// 	return nil, err
	// }

	// user.Username = username

	return nil, nil
}

func (r *UserRepository) GetHashPass(email string) (string, error) {
	// if email == "" {
	// 	return "", nil
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// var hashPass string
	// query := `SELECT hash_pass FROM users WHERE email = $1`

	// err := r.db.QueryRowContext(ctx, query, email).Scan(&hashPass)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return "", nil
	// 	}
	// 	return "", err
	// }

	return "", nil
}
