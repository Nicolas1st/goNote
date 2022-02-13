package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewNotesResourceRouter() {
	m := mux.NewRouter()

	m.HandleFunc("/notes/{id}", GetNote).Methods(http.MethodGet)
	m.HandleFunc("/notes", GetAllNotes).Methods(http.MethodGet)
	m.HandleFunc("/notes", CreateNote).Methods(http.MethodPost)
	m.HandleFunc("/notes/{id}", UpdateNote).Methods(http.MethodPut)
	m.HandleFunc("/notes/{id}", DeleteNote).Methods(http.MethodDelete)
}
