package main

import (
  "log"
  "os"
  "html/template"
)

type Page struct {
  Title string
  Heading string
  Input string
}

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){

  home := Page{
    Title: "Nothing is escaped",
    Heading: "Nothing is escaped with text/template",
    Input: `<script>alert("Yow!");</script>`,
  }

  err := tpl.Execute(os.Stdout, home)
  if err != nil {
    log.Fatalln(err)
    }

  // call functions in templates for formatting purposes only
  // not processing or data access - this should be done in
  // the controller (this script!) for seperation of concerns
}
