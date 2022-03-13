package api

import (
	"net/http"

	"github.com/Nicolas1st/goNote/app/model/entities/note"
	"github.com/gorilla/mux"
)

func NewNotesResourceRouter(n note.NoteEntity) *mux.Router {
	router := mux.NewRouter()
	router.Use(HeaderMiddleware)

	router.Handle("/{id:[0-9]+}/",
		BuildGet(n.GetOneByID),
	).Methods(http.MethodGet)

	router.Handle("/{author:[A-Za-z]+}/",
		BuildGetAll(n.GetAllByAuthor),
	).Methods(http.MethodGet)

	router.Handle("/",
		BuildCreate(n.StoreOne),
	).Methods(http.MethodPost)

	router.Handle("/{id:[0-9]+}/",
		BuildUpdate(n.UpdateOneByID),
	).Methods(http.MethodPut)

	router.Handle("/{id:[0-9]+}/",
		BuildDelete(n.DeleteOneByID),
	).Methods(http.MethodDelete)

	return router
}
