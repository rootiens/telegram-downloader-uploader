package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Schema() string {
	var create string = `CREATE TABLE IF NOT EXISTS files (
      id integer not null primary key,
      link text,
      fileid text,
      expire_at datetime);`
	return create
}

var DB *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	CheckErr(err)

	_, err = db.Exec(Schema())
	CheckErr(err)
	DB = db
	return DB
}

func GetDB() *sql.DB {
	return DB
}
