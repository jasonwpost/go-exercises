// serve an entire directory with http.FileServer
// and limit access with StripPrefix

package main

import (
  "io"
  "net/http"
)

func main(){
  http.HandleFunc("/", dog)
  // "/resources/" means all items beyond resources too
  // "/resources" would just be the folder itself not including any subfolders
  http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
  http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request){

  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w,`<img src="/assets/toby.jpg">`)
}
