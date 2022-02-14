package note

import "github.com/Nicolas1st/goNote/model/database"

type NoteEntity struct {
	DatabaseEnv *database.Env
}

func CreateNoteEntity(env *database.Env) *NoteEntity {
	return &NoteEntity{
		DatabaseEnv: env,
	}
}
