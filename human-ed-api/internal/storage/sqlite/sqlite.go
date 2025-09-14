package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func New(storagePath string) *Storage {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Fatalf("%s: cannot open connection for sqlite db: %s", op, err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("%s: ping failed: %s", op, err.Error())
	}

	return &Storage{DB: db}
}
