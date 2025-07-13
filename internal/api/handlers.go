package api

import (
	"encoding/json"
	"net/http"

	"github.com/Nicolas1st/goNote/internal/model"
	"github.com/gorilla/mux"
)

type NotesResource struct {
	repository NotesStore
}

func (resource *NotesResource) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok || id == "" {
		panic("invalid id")
	}

	note := resource.repository.GetNoteByID(id)
	json.NewEncoder(w).Encode(note)
}

func (resource *NotesResource) GetALlNotes(w http.ResponseWriter, r *http.Request) {
	// extract notes from the db
	notes, err := resource.repository.GetAllNotes()
	if err != nil {
		return
	}

	// send back the response
	json.NewEncoder(w).Encode(notes)
}

func (resource *NotesResource) PostNote(w http.ResponseWriter, r *http.Request) {
	note := &model.Note{}
	json.NewDecoder(r.Body).Decode(note)

	note, err := resource.repository.StoreNote(note)
	if err != nil {
		panic(err)
	}

	// send back the response
	json.NewEncoder(w).Encode(note)
}

func (resource *NotesResource) UpdateNoteByID(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok || id == "" {
		panic("invalid id")
	}

	var newNote model.Note
	
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		panic(err)
	}

	note := resource.repository.UpdateNoteByID(id, &newNote)

	json.NewEncoder(w).Encode(note)
}

func (resource *NotesResource) DeleteNoteByID(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		panic("invalid id")
	}

	note := resource.repository.DeleteNoteByID(id)

	json.NewEncoder(w).Encode(note)
}
