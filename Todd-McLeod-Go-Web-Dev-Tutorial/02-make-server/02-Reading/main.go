// Basic TCP server
// Three steps - listen, accept, read/write

package main

import (
  "fmt"
  "bufio"
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
    go handle(conn)
  }
}

func handle(conn net.Conn){
  scanner := bufio.NewScanner(conn)
  // each token is a line
  for scanner.Scan() {
    ln := scanner.Text()
    fmt.Println(ln)
  }
  defer conn.Close()

  // we never get here
  // we have an open stream connection
  fmt.Println("code got here")
}
