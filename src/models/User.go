package models

import (
	"errors"
	"strings"
	"time"
)

// using for first time omitempty
// when for pass this user to json and the id been in blanck, this not going to pass
// it don't going pass this for json, it go removed the id of json
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"Criadoem,omitempty"`
}

// method prepare is used to validate and format data that arrive in request
func (user *Usuario) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	user.formatter()
	return nil
}

func (user *Usuario) validate() error {
	if user.Nome == "" {
		return errors.New("Nome é obrigatório. Não pode ficar em branco")
	}

	if user.Nick == "" {
		return errors.New("Nick é obrigatório. Não pode ficar em branco")
	}

	if user.Email == "" {
		return errors.New("Email é obrigatório. Não pode ficar em branco")
	}

	if user.Senha == "" {
		return errors.New("Senha é obrigatório. Não pode ficar em branco")
	}

	// return value default of erro
	return nil
}

func (user *Usuario) formatter() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
