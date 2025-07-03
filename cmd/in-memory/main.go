package main

import (
	"fmt"

	"software-design-base/internal/storage"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("sqlite3", ":memory:?_foreign_keys=on")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`create table items (id integer primary key, value text)`)
	if err != nil {
		panic(err)
	}

	dbStorage := storage.New(db)

	dbStorage.Save("first string")
	dbStorage.Save("second string")
	fmt.Println("SAVED TO DB")

	fmt.Println(`RETRIEVING FROM DB:`)
	fmt.Println(dbStorage.Retrieve(1))
	fmt.Println(dbStorage.Retrieve(2))
}
