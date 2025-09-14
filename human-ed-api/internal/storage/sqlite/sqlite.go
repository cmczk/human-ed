package sqlite

import "database/sql"

type Storage struct {
	DB *sql.DB
}

func New() *Storage {
	return &Storage{}
}
