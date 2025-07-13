package main

import (
	"fmt"
	"net/http"

	"github.com/Nicolas1st/goNote/internal/api"
	"github.com/Nicolas1st/goNote/internal/database"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/")))

	// db initialization
	db, err := database.NewDatabase("sqlite3", "notes.db")
	if err != nil {
		panic(fmt.Errorf("Failed to initialized the database: %w", err))
	}
	defer db.CloseConnnection()

	http.Handle("/notes/", api.NewNotesResourceRouter(db))

	port := ":8880"
	fmt.Printf("Server is listening on port %v\n", port)
	http.ListenAndServe(port, nil)
}
