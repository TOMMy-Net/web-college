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
	if err != nil{
		return nil, err
	}
	CreateTables(db)
	return &Storage{DB: db}, nil
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

func (db *Storage) SaveUser(fn, ln, email string, password string)  error{
	b, _:= db.DB.Begin()
	data, err := db.DB.Prepare(`INSERT INTO accounts (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`)
	if err != nil{
		b.Rollback()
		return err
	}
	_, err = data.Exec(fn, ln, email, password)
	if err != nil{
		b.Rollback()
		return err
	}
	b.Commit()
	return nil
}

func (db *Storage) CheckUser(email string)  (bool, error){
	rows, err := db.DB.Query("SELECT email FROM accounts WHERE email = ?", email)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}


