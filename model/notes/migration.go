package notes

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (env *DBEnv) InitialMigration() {
	db, err := gorm.Open(
		env.DBDialect,
		env.DBName,
	)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Note{})
}
