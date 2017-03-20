package main

import (
  "fmt"
  "net/http"
)

func main(){
  http.HandleFunc("/", set)
  http.HandleFunc("/read", read)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request){
  http.SetCookie(w, &http.Cookie{
    Name: "my-cookie",
    Value: "some value",
  })
  fmt.Fprintln(w, "cookie written - check your browser")
}

func read(w http.ResponseWriter, req *http.Request){
  c, err := req.Cookie("my-cookie")
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  fmt.Fprintln(w, "your cookie:", c)
}
