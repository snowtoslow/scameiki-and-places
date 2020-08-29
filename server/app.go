package server

import (
	"database/sql"
	"net/http"
)

type App struct {
	httpServer *http.Server


}



func initDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil{
		return nil, err
	}

	if err := db.Ping(); err !=nil{
		return nil, err
	}

	return db, nil
}