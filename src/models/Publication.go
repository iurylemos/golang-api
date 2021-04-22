package models

import (
	"errors"
	"strings"
	"time"
)

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

// method prepare publication to validate and format data that arrive of request
func (publication *Publicacao) Prepare() error {
	if erro := publication.validate(); erro != nil {
		return erro
	}

	publication.formatter()

	return nil
}

func (publication *Publicacao) validate() error {
	if publication.Titulo == "" {
		return errors.New("titulo não pode ficar em branco")
	}

	if publication.Conteudo == "" {
		return errors.New("conteudo não pode ficar em branco")
	}

	return nil
}

func (publication *Publicacao) formatter() {
	publication.Titulo = strings.TrimSpace(publication.Titulo)
	publication.Conteudo = strings.TrimSpace(publication.Conteudo)
}
