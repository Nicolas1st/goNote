package notes

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	// ID        uint      //`json:"id" gorm:"column:id;type:INT NOT NULL;primaryKey;autoIncrement:true"`
	Title   string //`json:"title" gorm:"column:title;type:STRING NOT NULL"`
	Content string //`json:"content" gorm:"column:content;type:TEXT NOT NULL"`
	Author  string //`json:"author" gorm:"column:author;type:STRING NOT NULL"`
	// CreatedAt time.Time //`json:"ID" gorm:"column:id;type:INT NOT NULL;primaryKey;autoIncrement:true"`
	// UpdatedAt time.Time //`json:"ID" gorm:"column:id;type:INT NOT NULL;primaryKey;autoIncrement:true"`
}
