package repositories

import (
	"api-nos-golang/src/models"
	"database/sql"
	"fmt"
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

//Find users for nick or name
func (rep rep_users) Find(nameOrNick string) ([]models.Usuario, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	rows, erro := rep.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var users []models.Usuario

	for rows.Next() {
		var user models.Usuario
		// SCAN to insert values in user

		if erro = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		// insert user inside slice of users (append = acrescentar)
		users = append(users, user)

	}

	return users, nil
}

func (rep rep_users) FindID(id uint64) (models.Usuario, error) {
	rows, erro := rep.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", id)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer rows.Close()

	var user models.Usuario

	if rows.Next() {
		if erro = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return user, nil
}

func (rep rep_users) Update(id uint64, user models.Usuario) error {
	statement, erro := rep.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(user.Nome, user.Nick, user.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (rep rep_users) Delete(id uint64) error {
	statement, erro := rep.db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}
