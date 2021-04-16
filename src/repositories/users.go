package repositories

import (
	"api-nos-golang/src/models"
	"database/sql"
)

// this struct is used to received database
// the connection go be open in controller and pass to this struct
// and struct go make the interation with database
type rep_users struct {
	db *sql.DB
}

// create repository
func NewRepositoryUsers(db *sql.DB) *rep_users {
	return &rep_users{db}
}

// func Create insert user in database
// uint64 to no have values negatives
func (rep rep_users) Create(user models.Usuario) (uint64, error) {

	statement, erro := rep.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Nome, user.Nick, user.Email, user.Senha)

	if erro != nil {
		return 0, erro
	}

	lastIDInsert, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInsert), nil
}
