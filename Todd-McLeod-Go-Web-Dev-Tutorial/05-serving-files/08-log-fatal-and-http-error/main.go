// This is a cleaner way of serving a static webpage
// and will catch and log errors if/when they happen


package main

import (
  "net/http"
  "log"
)

func main(){
  log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

// http.Error allows you to specify the error given to the user
// i.e. 404 and a custom not found message
// example can be found in the serve-content directory
