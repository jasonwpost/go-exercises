package main

import (
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "net/http"
  "./controllers"
)

func main(){
  r := httprouter.New()
  uc := controllers.NewUserController(getSession())
  r.GET("/user/:id", uc.GetUser)
  r.POST("/user", uc.CreateUser)
  r.DELETE("/user/:id", uc.DeleteUser)
  http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
  // connect to local mongo
  s, err := mgo.Dial("mongodb://localhost")

  if err != nil {
    panic(err)
  }
  return s
}
