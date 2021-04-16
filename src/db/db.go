package db

import (
	"api-nos-golang/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// open connecting with database
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionDataBase)

	if erro != nil {
		return nil, erro
	}

	// check out the database
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	// if pass for theses errors, then open connect with database
	return db, nil
}
