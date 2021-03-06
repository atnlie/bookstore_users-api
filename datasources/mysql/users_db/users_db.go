package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ClientDb *sql.DB

	//use env instead
	userName     = os.Getenv("mysql_users_username")
	userPassword = os.Getenv("mysql_users_password")
	userHost     = os.Getenv("mysql_users_host")
	userDB       = os.Getenv("mysql_users_db_name")
)

func init() {
	fmt.Println("userName ", userName)
	databaseSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		userName,
		userPassword,
		userHost,
		userDB,
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
