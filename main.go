package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  router := mux.newRouter()
  log.Fatal(http.ListenAndServe(":8000", router))
}
