package main

import (
  "io"
  "net/http"
)

func main(){
  http.HandleFunc("/", foo)
  http.Handle("favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
  v := req.FormValue("q")
  io.WriteString(w, "Do my search: " + v)
}

// visit http://localhost:8080/?q=dog
