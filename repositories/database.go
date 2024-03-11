package repositories

import (
	"github.com/che-ict/DEV-DT-Microblog/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var connection *gorm.DB

func connect() {
	var err error
	connection, err = gorm.Open(sqlite.Open("db.sqlite"))

	if err != nil {
		log.Panic(err)
	}

	err = connection.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Panic(err)
	}
}

func DB() *gorm.DB {
	if connection == nil {
		connect()
	}
	return connection
}
