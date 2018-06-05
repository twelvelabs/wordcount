package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}

func DefaultEndpoint(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
}

func main() {
    fmt.Println("Starting the application...")

    r := mux.NewRouter()
    r.HandleFunc("/", DefaultEndpoint).Methods("GET")
    r.Use(LoggingMiddleware)

    // listen on port 80
    log.Fatal(http.ListenAndServe(":80", r))
}
