package note

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Author  string `json:"Author"`
}
