package main

import (
  "log"
  "os"
  "text/template"
  "time" // nanosecond precision.
  // should be used as values, not pointers. Value of type Time.
  // can be used by mulitple goroutines
  // has before, after, and equal methods
  // go == operator compares not just time but location also
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

// American style formatting
func monthDayYear(t time.Time) string {

  return t.Format("01-02-2006")
}

// European/NZ/AU/UK/literally-everywhere-else style formatting
func dayMonthYear(t time.Time) string {

  return t.Format("02-01-2006")
}

// {string - empty interface (implements everything)} pair
var fm = template.FuncMap{
  "fdateMDY": monthDayYear,
  "fdateDMY": dayMonthYear,

}
func main(){

  err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
  if err != nil {
    log.Fatalln(err)
    }
}
