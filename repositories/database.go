package repositories

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var _connection *gorm.DB

func _connect() {
	_connection, err := gorm.Open(sqlite.Open("db.sqlite"))

	if err != nil {
		log.Panic(err)
	}

	_connection.AutoMigrate()
}

func DB() *gorm.DB {
	if _connection == nil {
		_connect()
	}
	return _connection
}
