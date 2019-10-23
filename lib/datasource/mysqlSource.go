package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "J0h4n4_2019"
	dbName := "webroom"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Print("ok")
	}
	return db
}