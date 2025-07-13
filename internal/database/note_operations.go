package database

import (
	"fmt"

	"github.com/Nicolas1st/goNote/internal/model"
)

func (db *Database) GetNoteByID(id string) *model.Note {
	var note model.Note
	db.Where(&model.Note{NoteID: id}).First(&note)

	return &note
}

func (db *Database) GetAllNotesByAuthor(author string) *[]model.Note {
	var notes []model.Note
	db.Where(&model.Note{Author: author}).Find(&notes)

	return &notes
}

func (db *Database) StoreNote(note *model.Note) (*model.Note, error) {
	result := db.Create(note)
	if result.Error != nil {
		fmt.Printf("Error occured when saving a note: %v", result.Error)
	}

	return note, result.Error
}

func (db *Database) UpdateNoteByID(id string, fields *model.Note) *model.Note {
	var note model.Note
	db.Where(&model.Note{NoteID: id}).First(&note)

	note.Title = fields.Title
	note.Content = fields.Content
	db.Save(note)

	return &note
}

func (db *Database) DeleteNoteByID(id string) *model.Note {
	var note model.Note
	db.Unscoped().Where(&model.Note{NoteID: id}).Delete(&note)

	return &note
}
