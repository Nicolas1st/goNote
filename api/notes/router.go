package notes

import (
	"net/http"

	"github.com/Nicolas1st/goNote/api/middlewares"
	"github.com/gorilla/mux"
)

func NewNotesResourceRouter(repository NotesRepositoryInterface) *mux.Router {
	notesResource := &NotesResource{
		repository: repository,
	}

	router := mux.NewRouter()
	router.Use(middlewares.JsonResponseMiddleware)

	router.HandleFunc("/notes/{id:[0-9]+}/",
		notesResource.GetNoteByID,
	).Methods(http.MethodGet)

	router.HandleFunc("/notes/{author:[A-Za-z]+}/",
		notesResource.GetAllNotesByAuthor,
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
