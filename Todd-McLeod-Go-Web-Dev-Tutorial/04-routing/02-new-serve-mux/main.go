// Better solution - still not great

package main

import (
  "io"
  "net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request){
  io.WriteString(res, "doggy")
}

type hotcat int

func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request){
  io.WriteString(res, "kitty")
}

func main(){
  var d hotdog
  var c hotcat

  mux := http.NewServeMux()
  mux.Handle("/dog/", d) // also catches dog/whatever
  mux.Handle("/cat", c) // only catches /cat - /cat/something will not work
  http.ListenAndServe(":8080", mux)
}
