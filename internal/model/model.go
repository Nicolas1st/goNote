package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	NoteID      string `json:"NoteID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Author  string `json:"Author"`
}
