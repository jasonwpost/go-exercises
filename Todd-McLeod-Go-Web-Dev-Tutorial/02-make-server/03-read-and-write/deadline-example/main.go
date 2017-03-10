package main

import (
  "fmt"
  "bufio"
  "log"
  "net"
  "time"
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
  err := conn.SetDeadline(time.Now().Add(10 * time.Second))
  if err != nil {
    fmt.Fprintf(conn,"CONN TIMEOUT")
  }
  scanner := bufio.NewScanner(conn)
  // each token is a line
  for scanner.Scan() {
    ln := scanner.Text()
    fmt.Println(ln)
    fmt.Fprintf(conn, "I heard you say %s\n", ln)
  }
  defer conn.Close()

  fmt.Println("code gets here at timeout")
}
