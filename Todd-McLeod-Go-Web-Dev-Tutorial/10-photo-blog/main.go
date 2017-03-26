package main

import (
  "html/template"
  "net/http"
  "github.com/satori/go.uuid"
  "strings"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
  http.HandleFunc("/", index)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
  c := getCookie(w, req)
  c = assignValues(w, c);
  cookieSlice := strings.Split(c.Value, "|")
  tpl.ExecuteTemplate(w, "index.gohtml", cookieSlice)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
  c, err := req.Cookie("session")
  if err != nil {
    sID := uuid.NewV4()
    c = &http.Cookie {
      Name: "session",
      Value: sID.String(),
    }
    http.SetCookie(w, c)
  }
  return c;
}

func assignValues(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
  p1 := "cookies.jpg"
  p2 := "biscuits.jpg"
  p3 := "jaffacakes.jpg"
  str := c.Value
  if !strings.Contains(str, p1){
    str += "|" + p1
  }
  if !strings.Contains(str, p2){
    str += "|" + p2
  }
  if !strings.Contains(str, p3){
    str += "|" + p3
  }
  c.Value = str
  http.SetCookie(w, c)
  return c
}
