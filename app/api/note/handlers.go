package api

import (
	"encoding/json"
	"math/bits"
	"net/http"
	"strconv"

	"github.com/Nicolas1st/goNote/model/entities/note"
	"github.com/gorilla/mux"
)

func BuildGet(f func(id uint) *note.Note) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}

		note := *f(uint(id))

		json.NewEncoder(w).Encode(note)
	})

}

func BuildGetAll(f func(author string) *[]note.Note) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		author := vars["author"]
		notes := f(author)

		json.NewEncoder(w).Encode(notes)
	})
}

func BuildCreate(f func(*note.Note) (*note.Note, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		note := &note.Note{}
		json.NewDecoder(r.Body).Decode(note)

		note, err := f(note)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(note)
	})
}

func BuildUpdate(f func(id uint, title, content string) *note.Note) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}

		var noteWithFieldsToUpdate note.Note
		err = json.NewDecoder(r.Body).Decode(&noteWithFieldsToUpdate)
		if err != nil {
			panic(err)
		}

		note := f(uint(id), noteWithFieldsToUpdate.Title, noteWithFieldsToUpdate.Content)

		json.NewEncoder(w).Encode(note)
	})
}

func BuildDelete(f func(id uint) *note.Note) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, bits.UintSize)
		if err != nil {
			panic(err)
		}

		note := f(uint(id))

		json.NewEncoder(w).Encode(note)
	})
}
