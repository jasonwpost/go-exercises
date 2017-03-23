// turn into binary through this command
// GOOS=linux GOARCH=amd64 go build -o mybinary

package main

import (
  "net/http"
  "io"
)

func main(){
  http.HandleFunc("/", index)
  http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request){
  io.WriteString(w, "Oh yes, running on AWS.")
}
