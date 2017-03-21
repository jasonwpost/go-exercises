package main

import (
  "net/http"
  "github.com/satori/go.uuid"
  "html/template"
)

type user struct {
  UserName string
  First string
  Last string
}

var tpl *template.Template
var dbUsers = map[string]user{} // user id -> user
var dbSessions = map[string]string{} // session id -> user id
// could also declare as var dbSessions = make(map[string]string)

func init(){
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
  http.HandleFunc("/", foo)
  http.HandleFunc("/bar", bar)
  http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){

  // get cookie
  c, err := req.Cookie("session")
  if err != nil {
    sId := uuid.NewV4()
    c = &http.Cookie {
      Name: "session",
      Value: sId.String(),
      HttpOnly: true,
    }
  http.SetCookie(w, c)
  }

  // if user already exists, get user
  var u user
  if un, ok := dbSessions[c.Value]; ok {
    u = dbUsers[un]
  }

  // process form submission
  if req.Method == http.MethodPost {
    un := req.FormValue("username")
    f := req.FormValue("firstname")
    l := req.FormValue("lastname")
    u = user{un, f, l}
    dbSessions[c.Value] = un
    dbUsers[un] = u
  }

  tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request){
  // get cookie
  c, err := req.Cookie("session")
  if err != nil {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return;
  }
  un, ok := dbSessions[c.Value]

  if !ok {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  u := dbUsers[un]
  tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
