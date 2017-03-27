package main

import (
  "fmt"
  "log"
  "net/http"
  "context"
)

func main() {
  http.HandleFunc("/", foo)
  http.HandleFunc("/bar", bar)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

// should only add a few basics to context (user id, etc) -
//not designed to hold a lot of data
func foo(w http.ResponseWriter, req *http.Request){
  ctx := req.Context()

  ctx = context.WithValue(ctx, "userID", 777)
  ctx = context.WithValue(ctx, "fname", "Bond")

  results := dbAccess(ctx)

  fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
  uid := ctx.Value("userID").(int) ///.(int) asserts that this = an int value
  return uid
}

func bar(w http.ResponseWriter, req *http.Request){
  ctx := req.Context()

  log.Println(ctx)
  fmt.Fprintln(w, ctx)
}
