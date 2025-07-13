package api

import (
	"encoding/json"
	"math/bits"
	"net/http"
	"strconv"

	"github.com/Nicolas1st/goNote/internal/model"
	"github.com/gorilla/mux"
)

type NotesResource struct {
	repository NotesStore
}

func (resource *NotesResource) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, bits.UintSize)
	if err != nil {
		panic(err)
	}

	// extract the note from the db
	note := resource.repository.GetNoteByID(uint(id))

	// send back the response
	json.NewEncoder(w).Encode(note)
}

func (resource *NotesResource) GetAllNotesByAuthor(w http.ResponseWriter, r *http.Request) {
	// extract path param
	vars := mux.Vars(r)
	author := vars["author"]

	// extract notes from the db
	notes := resource.repository.GetAllNotesByAuthor(author)

	// send back the response
	json.NewEncoder(w).Encode(notes)
}

func (resource *NotesResource) PostNote(w http.ResponseWriter, r *http.Request) {
	// extracmodel. from the http request body
	note := &model.Note{}
	json.NewDecoder(r.Body).Decode(note)

	// store the note
	note, err := resource.repository.StoreNote(note)
	if err != nil {
		panic(err)
	}

	// send back the response
	json.NewEncoder(w).Encode(note)
}

func (resource *NotesResource) UpdateNoteByID(w http.ResponseWriter, r *http.Request) {
	// extract param from the url
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, bits.UintSize)
	if err != nil {
		panic(err)
	}

	// extract fiels from the http request body
	var noteWithFieldsToUpdate model.Note
	err = json.NewDecoder(r.Body).Decode(&noteWithFieldsToUpdate)
	if err != nil {
		panic(err)
	}

	// update the note in the db
	note := resource.repository.UpdateNoteByID(uint(id), &noteWithFieldsToUpdate)

	// send backe the response
	json.NewEncoder(w).Encode(note)
}

func (resource *NotesResource) DeleteNoteByID(w http.ResponseWriter, r *http.Request) {
	// extract param from the url
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, bits.UintSize)
	if err != nil {
		panic(err)
	}

	// delete the note in the db
	note := resource.repository.DeleteNoteByID(uint(id))

	// send send back the response
	json.NewEncoder(w).Encode(note)
}
