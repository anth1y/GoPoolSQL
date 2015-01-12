package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/lib/pq"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":5432")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
	//	server
	// have all incoming requests listen on multiple channels
	// then hand off to the postgres master
}
