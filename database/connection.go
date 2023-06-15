package database

import (
	"github.com/codingleo/auth-todo-backend/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Database struct {
	connectionString string
}

func NewConnection(connectionString string) *Database {
	return &Database{
		connectionString: connectionString,
	}
}

func (db *Database) Connect() {
	connection, err := gorm.Open(sqlite.Open(db.connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = connection.AutoMigrate(&types.User{})

	if err != nil {
		panic("failed to migrate database")
	}

	Db = connection
}
