package database

import (
	"database/sql"
	"log"
	"os"
	"pawpawchat/pkg/auth/database/repository"

	_ "github.com/lib/pq"
)

type DataBase struct {
	authInfo *repository.AuthInfoRepository
}

func New() *DataBase {
	dbDRIVER := os.Getenv("AUTH_DB_DRIVER")
	if dbDRIVER == "" {
		log.Fatal("missing db driver")
	}

	dbURL := os.Getenv("AUTH_DB_URL")
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

	return &DataBase{authInfo: repository.NewAuthInfoRepository(db)}
}

// Return UserAuth repository
func (d *DataBase) AuthInfo() *repository.AuthInfoRepository {
	return d.authInfo
}
