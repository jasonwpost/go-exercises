// Header interface :
// type Header interface {
//  Header() Header // returns the header map that will be sent by WriteHeader
//  Write([]byte) (int, error) // writes the data to the connect as part of an HTTP reply
//  WriteHeader(int) // sends an HTTP response header with status code
//}
// Must set Header before using WriteHeader

package main

import (
  "fmt"
  "net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
  w.Header().Set("Jason-Key", "this is from Jason")
  w.Header().Set("Content-Type", "text/html;charset=utf-8")
  fmt.Fprintln(w, "<h1>Any code you want in this function</h1>")
}

func main(){
  var d hotdog
  http.ListenAndServe(":8080", d)
}
