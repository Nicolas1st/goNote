package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewNotesResourceRouter(repository NotesStore) *mux.Router {
	notesResource := &NotesResource{
		repository: repository,
	}

	router := mux.NewRouter()
	router.Use(JsonResponseMiddleware)

	router.HandleFunc("/notes/{id:[0-9]+}/",
		notesResource.GetNoteByID,
	).Methods(http.MethodGet)

	router.HandleFunc("/notes/",
		notesResource.GetALlNotes,
	).Methods(http.MethodGet)

	router.HandleFunc("/notes/",
		notesResource.PostNote,
	).Methods(http.MethodPost)

	router.HandleFunc("/notes/{id:[0-9]+}/",
		notesResource.UpdateNoteByID,
	).Methods(http.MethodPut)

	router.HandleFunc("/notes/{id:[0-9]+}/",
		notesResource.DeleteNoteByID,
	).Methods(http.MethodDelete)

	return router
}
