package main

import(
  "encoding/base64"
  "fmt"
  "log"
)

func main(){
  s:= "Never memorize something that you can look up."

  s64 := base64.StdEncoding.EncodeToString([]byte(s))
  fmt.Println(s64)

  bs, err := base64.StdEncoding.DecodeString(s64) // byte slice
  if err != nil {
    log.Fatal("Error decoding")
  }
  fmt.Println(string(bs))
}
