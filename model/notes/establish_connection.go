package notes

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (env *DBEnv) connectToDB() (*gorm.DB, func()) {

	db, err := gorm.Open(env.DBDialect, env.DBName)
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect to the database")
	}
	fmt.Println("Connected to db")

	return db, func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("Db connection ws closed")
	}

}
