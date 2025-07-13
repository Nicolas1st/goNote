package database

import (
	"github.com/Nicolas1st/goNote/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(DatabaseDialect, DatabaseName string) (*Database, error) {
	connection, err := gorm.Open(DatabaseDialect, DatabaseName)
	database := &Database{connection}
	if err != nil {
		return database, err
	}

	database.performMigrations()

	return database, nil
}

func (db *Database) CloseConnnection() error {
	return db.Close()
}

func (db *Database) performMigrations() {
	db.AutoMigrate(&model.Note{})
}
