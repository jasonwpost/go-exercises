package main

import (
  "log"
  "os"
  "strings"
  "text/template"
)

var tpl *template.Template

type sage struct {
  Name string
  Motto string
}

type car struct {
  Manufacturer string
  Model string
  Doors int
}

// key is string, value is an empty interface with no methods
// therefore, every type implements this interface!
// uc is what the func will be called in the template
// ft is a func declared below
var fm = template.FuncMap{
  "uc": strings.ToUpper,
  "ft": firstThree,
}

func init(){
  // template.Must needs a pointer to a template, so must create a new
  // one, attach our fm functions (returns Template *), and then parse the files to
  // add our files
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
  s = strings.TrimSpace(s)
  s = s[:3]
  return s
}

func main(){

  b := sage{
    Name: "Buddha",
    Motto: "The belief of no beliefs",
  }

  g := sage{
    Name: "Gandhi",
    Motto: "Be the change",
  }

  m := sage{
    Name: "MLK",
    Motto: "Hatred never ceases with hatred but with love alone is healed.",
  }

  f := car {
  Manufacturer: "Ford",
  Model: "F150",
  Doors: 2,
  }

  c := car {
  Manufacturer: "Toyota",
  Model: "Corolla",
  Doors: 4,
  }

  sages := []sage{b, g, m}
  cars := []car{f, c}
  data := struct {
    Wisdom []sage
    Transport []car
  }{
    sages,
    cars,
}
  err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
  if err != nil {
    log.Fatalln(err)
    }
}
