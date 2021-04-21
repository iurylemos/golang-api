package models

import "time"

// field Curtidas were removed the "omitempty"
// why if i had one publication what not was "like" for no one
// the field "curtidas" would stay with value zero
// why it is the value default of uint64 is zero
// if i had passed the field "omitempty" for the field Curtidas, the field zero could not come

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autoNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}
