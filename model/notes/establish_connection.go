package notes

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (env *DBEnv) establishConnection(db *gorm.DB) func() {

	// hacky way to extract db connection code into a separate function
	// to avoid code duplication

	db, err := gorm.Open(env.DBDialect, env.DBName)
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect to the database")
	}

	return func() {
		db.Close()
	}

}
