package main

import (
	"fmt"
	"software-design-base/internal/storage"

	"github.com/jmoiron/sqlx"
)

var conn sqlx.Conn // abstract connect

func main() {
	dbStorage := storage.New(conn)
	dbStorage.Save("Data in DB")
	fmt.Println(dbStorage.Retrieve(0))
}
