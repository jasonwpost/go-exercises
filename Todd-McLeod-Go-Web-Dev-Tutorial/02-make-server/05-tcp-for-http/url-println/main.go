
package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
  "strings"
)

var sampString string;

func main(){
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln(err.Error())
  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Fatalln(err.Error())
    }
    go handle(conn)
  }
}

func handle(conn net.Conn){
  defer conn.Close()
  // read
  request(conn)
  // write
  respond(conn)
}

func request(conn net.Conn){
  i := 0
  scanner := bufio.NewScanner(conn)
  for scanner.Scan(){
    ln := scanner.Text()
    fmt.Println(ln)
    if i == 0 {
      // request
      m := strings.Fields(ln)[0]
      fmt.Println("***METHOD", m)
		}
    if (strings.Contains(ln, "Host")){
      sampString = ln[6:]
    }
		if ln == "" {
			// headers are done / terminating condition
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
  body = body + " " + sampString
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") // status
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) // header
	fmt.Fprint(conn, "Content-Type: text/html\r\n") // header
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
