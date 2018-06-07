package main

import (
    "encoding/json"
    "net/http"
)


type JsonError struct {
    Status  int     `json:"status"`
    Message string  `json:"message"`
}

type JwtToken struct {
    Token   string  `json:"token"`
}


// Render a struct as a JSON response
func RenderJson(w http.ResponseWriter, code int, obj interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(obj)
}
// Helper to render JSON errors (since the error already has the status code)
func RenderJsonError(w http.ResponseWriter, error JsonError) {
    RenderJson(w, error.Status, error)
}


func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
    var user User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        RenderJsonError(w, JsonError{ Status: http.StatusBadRequest, Message: "Invalid JSON" })
        return
    }

    _, err = NewUserService().AuthenticateCredentials(user.Name, user.Password)
    if err != nil {
        RenderJsonError(w, JsonError{ Status: http.StatusUnauthorized, Message: err.Error() })
        return
    }

    RenderJson(w, http.StatusOK, JwtToken{ Token: "lolwat" })
}
