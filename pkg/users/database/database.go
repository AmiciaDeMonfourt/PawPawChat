package database

import (
	"database/sql"
	"log"
	"os"
	"pawpawchat/pkg/users/database/repository"

	_ "github.com/lib/pq"
)

type DataBase struct {
	user *repository.UserRepository
}

func New() *DataBase {
	dbDRIVER := os.Getenv("USERS_DB_DRIVER")
	if dbDRIVER == "" {
		log.Fatal("missing db driver")
	}

	dbURL := os.Getenv("USERS_DB_URL")
	if dbURL == "" {
		log.Fatal("misssing db url")
	}

	db, err := sql.Open(dbDRIVER, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("could not connection to user data base: ", err)
	}

	return &DataBase{user: repository.NewUserRepository(db)}
}

func (d *DataBase) User() *repository.UserRepository {
	return d.user
}
