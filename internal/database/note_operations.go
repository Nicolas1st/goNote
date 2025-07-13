package database

import (
	"fmt"

	"github.com/Nicolas1st/goNote/internal/model"
)

func (db *Database) GetNoteByID(id uint) *model.Note {
	var note model.Note
	db.First(&note, id)

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
		fmt.Printf("Error occured: %v", result.Error)
	}

	return note, result.Error
}

func (db *Database) UpdateNoteByID(id uint, fields *model.Note) *model.Note {
	var note model.Note
	db.First(&note, id)

	note.Title = fields.Title
	note.Content = fields.Content
	db.Save(note)

	return &note
}

func (db *Database) DeleteNoteByID(id uint) *model.Note {
	var note model.Note
	db.Delete(&note, id)

	return &note
}
