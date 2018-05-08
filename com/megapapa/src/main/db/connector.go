package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	. "../entity"
	"fmt"
)
/*
	Now used MySQL
*/
func WriteParseResultToDb(parseResult *ParseResult) {
	fmt.Println(parseResult.Price)
	db, err := sql.Open("mysql", "root:password@/db")
	if err == nil {
		db.Exec("SELECT * FROM db.city")
	}
}
