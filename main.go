package main

import (
    "crypto/rsa"
    "log"
    "net/http"
    "io/ioutil"

    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/mux"
)

const sslCertPath       = "/home/wordcount/server.crt"
const sslKeyPath        = "/home/wordcount/server.key"

const jwtPrivateKeyPath = "/home/wordcount/jwt.key"
const jwtPublicKeyPath  = "/home/wordcount/jwt.key.pub"

var jwtPrivateKey = initJwtPrivateKey()
var jwtPublicKey  = initJwtPublicKey()

func initJwtPrivateKey() *rsa.PrivateKey {
    bytes, err := ioutil.ReadFile(jwtPrivateKeyPath)
    if err != nil {
        log.Fatal(err)
    }
    key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
    if err != nil {
        log.Fatal(err)
    }
    return key
}

func initJwtPublicKey() *rsa.PublicKey {
    bytes, err := ioutil.ReadFile(jwtPublicKeyPath)
    if err != nil {
        log.Fatal(err)
    }
    key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
    if err != nil {
        log.Fatal(err)
    }
    return key
}


func main() {
    log.Println("Starting the application...")

    r := mux.NewRouter()
    r.HandleFunc("/token", CreateTokenEndpoint).Methods("POST")
    r.HandleFunc("/wordcount", WordcountEndpoint).Methods("POST")
    r.Use(LoggingMiddleware)

    log.Fatal(http.ListenAndServeTLS(":443", sslCertPath, sslKeyPath, r))
}
