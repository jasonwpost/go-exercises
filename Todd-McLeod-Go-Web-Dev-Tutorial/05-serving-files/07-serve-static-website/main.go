// if you have a file called index.html, you can just
// serve a static website and the root dir will automatically
// show the index.html file

package main

import (
  "net/http"
)

func main(){
  http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
