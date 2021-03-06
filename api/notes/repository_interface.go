package notes

import "github.com/Nicolas1st/goNote/persistence/model"

type NotesRepositoryInterface interface {
	GetNoteByID(id uint) *model.Note
	GetAllNotesByAuthor(author string) *[]model.Note
	StoreNote(note *model.Note) (*model.Note, error)
	UpdateNoteByID(id uint, fields *model.Note) *model.Note
	DeleteNoteByID(id uint) *model.Note
}
