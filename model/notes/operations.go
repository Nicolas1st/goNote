package notes

import "github.com/jinzhu/gorm"

func (env *DBEnv) SelectNote(id uint) *Note {
	var db *gorm.DB
	defer env.establishConnection(db)()

	// db.Where(&Note{ID: id}).First(note)
	var note *Note
	db.First(note, id)

	return note
}

func (env *DBEnv) SelectNotesByAuthor(author string) *[]Note {
	var db *gorm.DB
	defer env.establishConnection(db)()

	var notes *[]Note
	db.Where(&Note{Author: author}).Find(notes)

	return notes
}

func (env *DBEnv) CreateNote(note *Note) (*Note, error) {
	var db *gorm.DB
	defer env.establishConnection(db)()

	result := db.Create(note)

	return note, result.Error
}

func (env *DBEnv) UpdateNote(id uint, title, content string) *Note {
	var db *gorm.DB
	defer env.establishConnection(db)()

	var note *Note
	db.First(note, id)
	// db.Where(&Note{ID: id}).First(note)

	note.Title = title
	note.Content = content
	db.Save(note)

	return note
}

func (env *DBEnv) DeleteNote(id uint) *Note {
	var db *gorm.DB
	defer env.establishConnection(db)()

	var note *Note
	db.Delete(note, id)
	// db.Where(&Note{ID: id}).Delete(note)

	return note
}
