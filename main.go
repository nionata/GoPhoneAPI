package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

var people []Person

func main() {
  fmt.Printf("API server started on port 8000")

  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}
