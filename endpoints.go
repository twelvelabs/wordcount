package main

import (
    "net/http"
)

func DefaultEndpoint(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
}
