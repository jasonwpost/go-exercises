package main

import (
  "log"
  "os"
  "text/template"
)

type person struct {
  Name string
  Age int
}

func (p person) SomeProcessing() int {
  return 7
}

func (p person) AgeDbl() int {
  return p.Age * 2
}

func (p person) TakesArgs(x int) int {
  return x * 2
}

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){

  p := person {
    Name: "Ian Fleming",
    Age: 56,
}

  err := tpl.Execute(os.Stdout, p)
  if err != nil {
    log.Fatalln(err)
    }

  // call functions in templates for formatting purposes only
  // not processing or data access - this should be done in
  // the controller (this script!) for seperation of concerns
}
