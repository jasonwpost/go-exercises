// Run 02-Reading/main.go to serve this

package main

import (
  "fmt"
  "net"
)

func main(){
  conn, err := net.Dial("tcp", ":8080")
  if err != nil {
    panic(err)
  }
  defer conn.Close()

  fmt.Fprintln(conn, "I dialed you.")
}
