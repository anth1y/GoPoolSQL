package lb

import (
	log "github.com/Sirupsen/logrus"
	"io"
	"net"
)

func lb() {
	// Let's listen on tcp port 8149 on all interfaces
	// Make a struct for each port interface 8149 5432 and structs for ReadOnly ReadWrite
	// Lastly we will make channels and make them concurrent
	l, err := net.Listen("tcp", ":8149")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle each connection in a new goroutine
		go func(c net.Conn) {
			// not sure what function to call here io.
			// We want to receive sql msgs so maybe a io.Read or io.write?
			// Shut down the connect
			c.Close()
		}(conn)
	}
}
