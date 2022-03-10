package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (env *Env) ConnectToDatabase() (*gorm.DB, func()) {
	db, err := gorm.Open(env.DatabaseDialect, env.DatabaseName)
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect to the database")
	}

	return db, func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}
}
