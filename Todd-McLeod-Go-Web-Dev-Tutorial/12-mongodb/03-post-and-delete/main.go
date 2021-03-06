package main

import (
  "encoding/json"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "./models"
)

func main(){
  r := httprouter.New()
  r.GET("/", index)
  r.GET("/user/:id", getUser)
  r.POST("/user", createUser)
  r.DELETE("/user/:id", deleteUser)
  http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
  s := `<!DOCTYPE html>
          <html lang="en">
            <head>
              <meta charset="UTF-8">
              <title>Index</title>
            </head>
            <body>
              <a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
            </body>
          </html>
	`
  w.Header().Set("Content-Type", "text/html;charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  u := models.User{
    Name: "James Bond",
    Gender: "male",
    Age: 32,
    Id: p.ByName("id"),
  }

  // Marshal into json
  uj, err := json.Marshal(u)
  if err != nil {
    fmt.Println(err)
  }

  // Write content-type, status code, payload
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  u := models.User{}

  json.NewDecoder(r.Body).Decode(&u)

  u.Id = "007"

  uj, _ := json.Marshal(u)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "Write code to delete user\n")
}
