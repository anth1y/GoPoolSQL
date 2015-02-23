package main

import (
  "io"
  log "github.com/Sirupsen/logrus"
  "net"
)

func main() {
   // Let's listen on tcp port 8149 on all interfaces
   l, err := net.Listen("tcp", ":8149")
   if err := nil {
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
    } (conn)
  }
}
