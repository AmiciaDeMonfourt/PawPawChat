package database

import (
	"log"
	"os"
	"pawpawchat/pkg/auth/database/repository"
	"pawpawchat/pkg/auth/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	authInfo *repository.AuthInfoRepository
}

func New() *DataBase {
	dbURL := os.Getenv("AUTH_DB_URL")
	if dbURL == "" {
		log.Fatal("misssing db url")
	}

	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.AuthInfo{}); err != nil {
		log.Fatal(err)
	}

	return &DataBase{authInfo: repository.NewAuthInfoRepository(db)}
}

// Return UserAuth repository
func (d *DataBase) AuthInfo() *repository.AuthInfoRepository {
	return d.authInfo
}
