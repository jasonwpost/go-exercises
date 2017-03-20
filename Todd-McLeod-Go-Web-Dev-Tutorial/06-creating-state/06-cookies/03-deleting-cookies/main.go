// Cookies usually have a limit of 4096 chars
// chrome currently is set to a max of 180 Cookies
// written per site. For other browsers including IE
// it is much lower. Safari's limit is around 600

package main

import (
  "fmt"
	"net/http"
)

func main() {
  http.HandleFunc("/", set)
  http.HandleFunc("/read", read)
  http.HandleFunc("/expire", expire)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
  http.SetCookie(w, &http.Cookie{
    Name:  "session",
    Value: "some value",
  })
  fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {

  c, err := req.Cookie("session")
	if err != nil {
    http.Redirect(w, req, "/set", http.StatusSeeOther)
    return
	}
    fmt.Fprintln(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
  c, err := req.Cookie("session")
  if err != nil {
    http.Redirect(w, req, "/set", http.StatusSeeOther)
    return
  }
  c.MaxAge = -1 // this deletes the cookie
  http.SetCookie(w, c)
  http.Redirect(w, req, "/", http.StatusSeeOther)
}
