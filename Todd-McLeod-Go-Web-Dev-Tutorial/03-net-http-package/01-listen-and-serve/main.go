// Handler :
// type Handler interface {
//  ServeHTTP(ResponseWriter, *Request)
//}

// ListenAndServe basically does everything explained
// in the 02-make-server folder automagically.

package main

import (
  "fmt"
  "net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "Any code you want in this func")
}

func main(){
  var d hotdog
  http.ListenAndServe(":8080", d)
}
