package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
        "database/sql"
	_ "github.com/lib/pq"
        "time"
)

const (
  gophers = 10
  entries = 10000
)

func main() {
        // create a statement string
      var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
      
         // run the insert function using 10 go routines
        for i := 0; i < gophers; i++{
         // spin up a gopher
        go gopher(i, sStmt)
        }
    
        // this is a simple way to keep a program open
        // the go program will close when a key is pressed
        var input string
        fmt.Scanln(&input)
}
 
func gopher(gopher_id int, sStmt string){
        //This is to open the db things a bit but will not open till first req. 
      db, err := sql.Open("postgres",  "host=/var/run/postgresql dbname=testdb user=ant sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

     fmt.Printf("Gopher Id: %v || StartTime: %v\n", gopher_id, time.Now())

     for i := 0; i < entries; i++ { 

       stmt, err := db.Prepare(sStmt)
	if err != nil {
	  log.Fatal(err)
	}
      
        // close statement
        stmt.Close()
    } 
    // close db
    db.Close()
         
   fmt.Printf("Gopher Id: %v || StopTime: %v\n", gopher_id, time.Now())

}

