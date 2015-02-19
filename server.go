package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
        "database/sql"
	_ "github.com/lib/pq"
        "time"
)

func main() {
        // create a statement string
        var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
        //This is to open the db things a bit but will not open till first req. 
	db, err := sql.Open("postgres",  "host=/var/run/postgresql dbname=testdb user=ant sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

       stmt, err := db.Prepare(sStmt)
	if err != nil {
		log.Fatal(err)
	}

       fmt.Printf("StartTime: %v\n", time.Now())

//       res, err := stmt.Exec(1, time.Now())
//	if err != nil {
//		log.Fatal(err)
//	}

        // close statement
        stmt.Close()
        
        // close db
        db.Close()
         
       fmt.Printf("StopTime: %v\n", time.Now())
}
