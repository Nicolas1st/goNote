package api

import (
	"github.com/Nicolas1st/goNote/internal/model"
)

type NotesStore interface {
	GetNoteByID(id string) *model.Note
	GetAllNotes() (*[]model.Note, error)
	StoreNote(note *model.Note) (*model.Note, error)
	UpdateNoteByID(id string, fields *model.Note) *model.Note
	DeleteNoteByID(id string) *model.Note
}
