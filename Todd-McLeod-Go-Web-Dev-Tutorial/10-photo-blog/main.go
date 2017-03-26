package main

import (
  "html/template"
  "net/http"
  "github.com/satori/go.uuid"
  "strings"
  "fmt"
  "io"
  "crypto/sha1"
  "os"
  "path/filepath"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
  http.HandleFunc("/", index)
  http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
  c := getCookie(w, req)
  if req.Method == http.MethodPost {
    mf, fh, err := req.FormFile("nf")
    if err != nil {
      fmt.Println(err)
    }
    defer mf.Close()
    // create sha1 for file Name
    ext := strings.Split(fh.Filename, ".")[1]
    h := sha1.New()
    io.Copy(h, mf)
    fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
    // create new file
    wd, err := os.Getwd() // get current working dir
    if err != nil {
      fmt.Println(err)
    }
    path := filepath.Join(wd, "public", "pics", fname)
    nf, err := os.Create(path)
    if err != nil {
      fmt.Println(err)
    }
    defer nf.Close()
    // copy
    mf.Seek(0, 0)
    io.Copy(nf, mf) // copy into new file nf from mf
    // add file names to this user
    c = assignValues(w, c, fname);
  }
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

func assignValues(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
  str := c.Value
  if !strings.Contains(str, fname){
    str += "|" + fname
  }
  c.Value = str
  http.SetCookie(w, c)
  return c
}
