package main

import (
	"database/sql"
	"log"
	"mailinglist/jsonapi"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DbPath = "list.db"
	Port   = ":8080"
)

func main() {
	log.Printf("using database '%v'\n", DbPath)
	db, err := sql.Open("sqlite3", DbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Printf("starting JSON API server...\n")
	jsonapi.Serve(db, Port)
}
