package main

import (
  "html/template"
  "net/http"
  "fmt"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
  http.HandleFunc("/", index)
  http.HandleFunc("/foo", foo)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
  s := `here is some text from foo`
  fmt.Fprintln(w, s)
}
