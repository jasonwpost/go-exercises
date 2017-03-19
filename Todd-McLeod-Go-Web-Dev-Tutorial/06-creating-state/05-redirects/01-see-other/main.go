package main

import (
  "fmt"
  "html/template"
  "net/http"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
  http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

// redirecting function
func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form submission here
  // other redirect possibilties are
  // http.StatusTemporaryRedirect - keeps method the same
  // http.StatusMovedPermanately - browser will cache new location

  // could also use instead of http.Redirect
  // w.Header().Set("Location", "/")
  // w.WriterHeader(http.StatusSeeOther)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
