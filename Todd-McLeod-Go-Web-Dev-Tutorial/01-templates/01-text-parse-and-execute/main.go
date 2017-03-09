package main

import (
  "log"
  "os"
  "text/template"
)

func main(){

  // Two steps - parse files, then execute files


  // returns pointer to template object
  // tpl is a container holding all files are are parsed
  tpl, err := template.ParseFiles("tpl.gohtml")
  if err != nil {
    log.Fatalln(err)
  }

  nf, err := os.Create("index.html")
  if err != nil {
    log.Println("error creating file", err)
  }
  defer nf.Close()

  // template.Execute takes writer and reader
  err = tpl.Execute(nf, nil)
  if err != nil {
    log.Fatalln(err)
  }

  // Alternative to template.Execute is template.ExecuteTemplate
  // This executes a specific template and takes arguments
  // (writer, name, reader). The template which name refers
  // to is a template added to a container through ParseFiles

  // Another option is template.ParseGlob which matches a pattern
  // specified in the arugments
  // ex. tempalte.ParseGlob("templates/*.gohtml") for all files
  // of type gohtml in the templates subfolder

  // We can also use template.Must to do our error checking for us
  // must takes a pointer to a template and an error (which the Parse
  // functions all return), and then Must returns just the template
  // pointer
}
