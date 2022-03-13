package main

import (
	"fmt"
	"net/http"

	api "github.com/Nicolas1st/goNote/app/api/note"
	"github.com/Nicolas1st/goNote/app/model/database"
	entity "github.com/Nicolas1st/goNote/app/model/entities/note"
	"github.com/Nicolas1st/goNote/app/model/migrations"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	// serving the index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/html/index.html")
	})

	// serving static files
	fs := http.FileServer(http.Dir("../static/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", fs))

	// db initialization
	dbEnv := database.Env{
		DatabaseDialect: "sqlite3",
		DatabaseName:    "test.db",
	}

	// creating tables in the db if the don't exist
	migrations.Initialize(&dbEnv)

	// creating a router for notes resource
	notesResource := api.NewNotesResourceRouter(entity.NoteEntity{
		DatabaseEnv: &dbEnv,
	})

	http.Handle("/notes", notesResource)

	// running the server
	port := ":8880"
	fmt.Printf("Server is listening on port %v\n", port)
	http.ListenAndServe(port, nil)

}
