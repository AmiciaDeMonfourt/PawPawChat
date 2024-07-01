package db

import (
	"context"
	"database/sql"
	"pawpawchat/pkg/users/config"
	"pawpawchat/pkg/users/database/repository"

	_ "github.com/lib/pq"
)

type DataBase struct {
	user *repository.UserRepository
}

func New() (*DataBase, error) {
	cfg := config.DataBase()

	db, err := sql.Open(cfg.Driver, cfg.Source)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(context.TODO())
	if err != nil {
		return nil, err
	}

	return &DataBase{
		user: repository.NewUserRepository(db),
	}, nil
}

func (d *DataBase) User() *repository.UserRepository {
	return d.user
}
