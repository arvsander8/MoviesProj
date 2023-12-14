package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (App *Application) ConnectToDB() (*sql.DB, error) {
	connection, err := openDB(App.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")
	return connection, nil
}
