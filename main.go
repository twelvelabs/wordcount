package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    fmt.Println("Starting the application...")

    r := mux.NewRouter()
    r.HandleFunc("/", DefaultEndpoint).Methods("GET")
    r.Use(LoggingMiddleware)

    // listen on port 80
    log.Fatal(http.ListenAndServe(":80", r))
}
