package note

import "fmt"

func (n *NoteEntity) GetOneByID(id uint) *Note {
	db, closeConnection := n.DatabaseEnv.ConnectToDatabase()
	defer closeConnection()

	var note Note
	db.First(&note, id)

	return &note
}

func (n *NoteEntity) GetAllByAuthor(author string) *[]Note {
	db, closeConnection := n.DatabaseEnv.ConnectToDatabase()
	defer closeConnection()

	var notes []Note
	db.Where(&Note{Author: author}).Find(&notes)

	return &notes
}

func (n *NoteEntity) StoreOne(note *Note) (*Note, error) {
	db, closeConnection := n.DatabaseEnv.ConnectToDatabase()
	defer closeConnection()

	result := db.Create(note)
	if result.Error != nil {
		fmt.Printf("Error occured: %v", result.Error)
	}

	return note, result.Error
}

func (n *NoteEntity) UpdateOneByID(id uint, title, content string) *Note {
	db, closeConnection := n.DatabaseEnv.ConnectToDatabase()
	defer closeConnection()

	var note Note
	db.First(&note, id)

	note.Title = title
	note.Content = content
	db.Save(note)

	return &note
}

func (n *NoteEntity) DeleteOneByID(id uint) *Note {
	db, closeConnection := n.DatabaseEnv.ConnectToDatabase()
	defer closeConnection()

	var note Note
	db.Delete(&note, id)

	return &note
}
