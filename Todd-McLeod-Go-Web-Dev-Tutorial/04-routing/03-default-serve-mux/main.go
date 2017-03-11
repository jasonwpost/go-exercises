// more elegant solution
// giving ListenAndServe a nil handler
// means it will use the default ServeMux
// found in the standard library.
// We can then just assign functions to http.HandleFunc

package main

import (
  "io"
  "net/http"
)

func d(m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request){
  io.WriteString(res, "doggy")
}

func c(m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request){
  io.WriteString(res, "kitty")
}

func main(){

  http.HandleFunc("/dog/", d) // also catches dog/whatever
  http.HandleFunc("/cat", c) // only catches /cat - /cat/something will not work

  http.ListenAndServe(":8080", nil)
}
