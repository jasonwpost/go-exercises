// Basic TCP server

package main

import (
  "fmt"
  "io"
  "log"
  "net"
)

func main(){
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Panic(err)
  }
  defer li.Close()

  for {
    //connection implements writer and reader interfaces
    conn, err := li.Accept()
    if err != nil {
      log.Println(err)
    }
    // WriteString needs a writer and a string
    // connection implements the writer inferface
    io.WriteString(conn, "\nHello from TCP server\n")
    fmt.Fprintln(conn, "How is your day?")
    fmt.Fprintf(conn, "%v", "Well, I hope!")

    conn.Close()
  }
}
