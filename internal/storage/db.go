package storage

import (
	"github.com/jmoiron/sqlx"
)

type dbStorage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *dbStorage {
	return &dbStorage{
		db: db,
	}
}

func (ds *dbStorage) Save(value string) {
	ds.db.DB.Exec(
		`insert into items (value) values ($1)`,
		value,
	)
}

func (ds *dbStorage) Retrieve(id int) string {
	values := []string{}

	err := ds.db.Select(
		&values,
		`select value from items where id=$1`,
		id,
	)

	if err != nil || len(values) == 0 {
		return ""
	}

	return values[0]
}
