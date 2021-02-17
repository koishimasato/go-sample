package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
)

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	UserType int    `db:"user_type"`
}

var dbmap *gorp.DbMap

func init() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3314)/sample_db")
	checkErr(err, "sql.Open failed")

	_dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	_dbmap.AddTableWithName(User{}, "users").SetKeys(false, "id")

	err = _dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	dbmap = _dbmap
}

func GetDBMap() *gorp.DbMap {
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}