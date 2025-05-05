package db

import (
	"database/sql"
	_ "embed"
	"github.com/jettdc/cortex/utils"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//go:embed schema.sql
var schemaFile string

var database *sql.DB

func InitDb() {
	var err error
	database, err = sql.Open("sqlite3", utils.GetCortexPath("cortex.db"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec(schemaFile)
	if err != nil {
		log.Fatal(err)
	}
}
