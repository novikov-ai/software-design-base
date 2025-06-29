package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type dbStorage struct {
	conn sqlx.Conn
}

func New(c sqlx.Conn) *dbStorage {
	return &dbStorage{
		conn: c,
	}
}

func (ds *dbStorage) Save(value string) {
	ds.conn.ExecContext(
		context.Background(),
		`insert into items (value) values ($1)`,
		value,
	)
}

func (ds *dbStorage) Retrieve(id int) string {
	res, err := ds.conn.ExecContext(
		context.Background(),
		`select from items where id=$1`,
		id,
	)

	if err != nil {
		return ""
	}

	insertedID, err := res.LastInsertId()
	if err != nil || insertedID == 0 {
		return ""
	}

	return fmt.Sprintf("%d", insertedID)
}
