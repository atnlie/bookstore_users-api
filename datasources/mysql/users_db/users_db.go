package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ClientDb *sql.DB
)

func init() {
	databaseSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		"atnlie",
		"Admin12345",
		"127.0.0.1:3306",
		"users_db",
	)

	var err error
	ClientDb, err = sql.Open("mysql", databaseSource)
	if err != nil {
		panic(err)
	}
	if err = ClientDb.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
}