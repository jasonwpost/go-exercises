package controllers

import (
  "encoding/json"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "net/http"
  "../models"
)

// so our controller has access to a mongo session
type UserController struct {
  session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
  return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
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

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

  u := models.User{}

  json.NewDecoder(r.Body).Decode(&u)

  u.Id = "007"

  uj, _ := json.Marshal(u)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "Write code to delete user\n")
}
