package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}
func DBLite() (*Storage, error){
	db, err := sql.Open("sqlite3", "database/sqlite.db")
	
	CreateTables(db)
	return &Storage{DB: db}, err
}

func CreateTables(db *sql.DB)  {
	data, err := db.Prepare(`CREATE TABLE IF NOT EXISTS accounts (
		first_name TEXT NOT NULL,
		last_name  TEXT NOT NULL,
		email      TEXT NOT NULL UNIQUE,
		password   TEXT NOT NULL
		);`)
	if err != nil {
		log.Fatal(err)
	}
	data.Exec()
}

func (db *Storage) SaveUser(fn, ln, email, password string)  error{
	data, err := db.DB.Prepare(`INSERT INTO accounts (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`)
	if err != nil{
		return err
	}
	_, err = data.Exec(fn, ln, email, password)
	if err != nil{
		return err
	}
	return nil
}
