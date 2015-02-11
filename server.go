package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
        "database/sql"
	_ "github.com/lib/pq"
	"net"
)

func main() {
        // create a statement string
        var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
        //This is to open the db things a bit 
	db, err := sql.Open("postgres",  "user=pachygo dbname=pachygo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

       stmt, err := db.Prepare(sStmt)
	if err != nil {
		log.Fatal(err)
	}

       fmt.Printf("StartTime: %v\n", time.Now())

       res, err := stmt.Exec(1, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	// Write a single query to insert into a db.
	//	server
	// have all incoming requests listen on multiple channels
	// then hand off to the postgres master
}
