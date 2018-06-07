package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

const sslCertPath string = "/home/wordcount/server.crt"
const sslKeyPath string = "/home/wordcount/server.key"

func main() {
    fmt.Println("Starting the application...")

    r := mux.NewRouter()
    r.HandleFunc("/token", CreateTokenEndpoint).Methods("POST")
    r.Use(LoggingMiddleware)

    log.Fatal(http.ListenAndServeTLS(":443", sslCertPath, sslKeyPath, r))
}
