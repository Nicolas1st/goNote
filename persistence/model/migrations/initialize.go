package migrations

import (
	"github.com/Nicolas1st/goNote/persistence/model/database"
	"github.com/Nicolas1st/goNote/persistence/model/entities/note"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Initialize(env *database.Env) {
	db, closeConnection := env.ConnectToDatabase()
	defer closeConnection()

	db.AutoMigrate(&note.Note{})
}
