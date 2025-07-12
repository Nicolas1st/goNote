package main

import (
	"fmt"
	"net/http"

	"github.com/Nicolas1st/goNote/api/auth"
	"github.com/Nicolas1st/goNote/api/middlewares"
	"github.com/Nicolas1st/goNote/api/notes"
	"github.com/Nicolas1st/goNote/api/views"
	"github.com/Nicolas1st/goNote/persistence/database"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// the reason for serving them separately is that
	// there are other files in directories used for development
	// it's temporary

	// serving html
	viewsRouter := views.NewViewsRouter("./web/html")
	http.Handle("/", middlewares.AuthenticationMiddleware(viewsRouter, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You need to login"))
	}))

	// adding authentication
	http.Handle("/auth/", auth.NewAuthRouter())

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./web/html/index.html")
	// })

	// serving css
	css := http.FileServer(http.Dir("./web/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", css))

	// serving js
	js := http.FileServer(http.Dir("./web/js"))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", js))

	// db initialization
	db, err := database.NewDatabase("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	// creating a router for notes resource
	notesRouter := notes.NewNotesResourceRouter(db)
	http.Handle("/notes/", notesRouter)

	// running the server
	port := ":8880"
	fmt.Printf("Server is listening on port %v\n", port)
	http.ListenAndServe(port, nil)
}
