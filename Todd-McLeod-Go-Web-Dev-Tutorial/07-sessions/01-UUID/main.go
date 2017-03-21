package main

import (
  "fmt"
  "github.com/satori/go.uuid"
  "net/http"
)

func main(){
  http.HandleFunc("/", foo)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
  cookie, err := req.Cookie("session-id")
  if err != nil {
    id := uuid.NewV4()
    cookie = &http.Cookie {
      Name: "session-id",
      Value: id.String(),
      //Secure: true // this means it will only work over https
      HttpOnly: true, // this means you can't access the cookie with javascript
    }
  http.SetCookie(w, cookie)
  }
  fmt.Println(cookie)
}
