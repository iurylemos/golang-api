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

func (rep rep_publications) FindForID(id uint64) (models.Publicacao, error) {
	rows, erro := rep.db.Query(`
		SELECT p.*, u.nick FROM
		publicacoes p INNER JOIN usuarios u 
		ON u.id = p.autor_id WHERE p.id = ?
	`, id)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer rows.Close()

	var publication models.Publicacao

	if rows.Next() {
		if erro = rows.Scan(
			&publication.ID,
			&publication.Titulo,
			&publication.Conteudo,
			&publication.AutorID,
			&publication.Curtidas,
			&publication.CriadaEm,
			&publication.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publication, nil
}

func (rep rep_publications) FindPublications(id uint64) ([]models.Publicacao, error) {
	rows, erro := rep.db.Query(`
		SELECT DISTINCT p.*, u.nick FROM publicacoes p
		INNER JOIN usuarios u ON u.id = p.autor_id
		INNER JOIN seguidores s ON p.autor_id = s.usuario_id
		WHERE u.id = ? OR s.seguidor_id = ?
		ORDER BY 1 desc
	`, id, id)

	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var publications []models.Publicacao

	for rows.Next() {
		var publication models.Publicacao

		if erro = rows.Scan(
			&publication.ID,
			&publication.Titulo,
			&publication.Conteudo,
			&publication.AutorID,
			&publication.Curtidas,
			&publication.CriadaEm,
			&publication.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (rep rep_publications) Update(id uint64, publication models.Publicacao) error {
	statement, erro := rep.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(publication.Titulo, publication.Conteudo, id); erro != nil {
		return erro
	}

	return nil
}

func (rep rep_publications) Delete(id uint64) error {
	statement, erro := rep.db.Prepare("DELETE FROM publicacoes WHERE id = ?")

	if erro != nil {
		return nil
	}

	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (rep rep_publications) FindPublicationsForUser(id uint64) ([]models.Publicacao, error) {
	rows, erro := rep.db.Query(`
		SELECT p.*, u.nick FROM publicacoes p
		JOIN usuarios u ON u.id = p.autor_id
		WHERE p.autor_id = ?
	`, id)

	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var publications []models.Publicacao

	for rows.Next() {
		var publication models.Publicacao

		if erro = rows.Scan(
			&publication.ID,
			&publication.Titulo,
			&publication.Conteudo,
			&publication.AutorID,
			&publication.Curtidas,
			&publication.CriadaEm,
			&publication.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}
