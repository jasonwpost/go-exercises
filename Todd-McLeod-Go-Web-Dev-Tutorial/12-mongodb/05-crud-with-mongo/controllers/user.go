package controllers

import (
  "encoding/json"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
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
  id := p.ByName("id")

  if !bson.IsObjectIdHex(id){
    w.WriteHeader(http.StatusNotFound)
    return
  }

  oid := bson.ObjectIdHex(id)

  u := models.User{}

  if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u);err != nil {
    w.WriteHeader(404)
    return
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

  u.Id = bson.NewObjectId()

  uc.session.DB("go-web-dev-db").C("users").Insert(u)

  uj, err := json.Marshal(u)
  if err != nil {
    fmt.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  id := p.ByName("id")

  if !bson.IsObjectIdHex(id){
    w.WriteHeader(404)
    return
  }

  oid := bson.ObjectIdHex(id)

  if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
    w.WriteHeader(404)
    return
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "Deleted user", oid, "\n")
}
