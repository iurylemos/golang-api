package repositories

import (
	"api-nos-golang/src/models"
	"database/sql"
)

// this struct is used to received database
// the connection go be open in controller and pass to this struct
// and struct go make the interation with database
type rep_publications struct {
	db *sql.DB
}

// create repository
func NewRepositoryPublications(db *sql.DB) *rep_publications {
	return &rep_publications{db}
}

func (rep rep_publications) Create(publication models.Publicacao) (uint64, error) {
	statament, erro := rep.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statament.Close()

	result, erro := statament.Exec(publication.Titulo, publication.Conteudo, publication.AutorID)

	if erro != nil {
		return 0, erro
	}

	lastIDInsert, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInsert), nil

}
