package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewNotesResourceRouter() {
	m := mux.NewRouter()

	m.HandleFunc("/{id}", GetNote).Methods(http.MethodGet)
	m.HandleFunc("/", GetAllNotes).Methods(http.MethodGet)
	m.HandleFunc("/", CreateNote).Methods(http.MethodPost)
	m.HandleFunc("/{id}", UpdateNote).Methods(http.MethodPut)
	m.HandleFunc("/{id}", DeleteNote).Methods(http.MethodDelete)
}
