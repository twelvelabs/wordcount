package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
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
        RenderJsonError(w, JsonError{ Status: http.StatusUnauthorized, Message: "Invalid credentials" })
        return
    }

    // Create the Claims
    claims := &jwt.StandardClaims{
        IssuedAt:   time.Now().Unix(),
        ExpiresAt:  time.Now().Add(time.Minute * time.Duration(5)).Unix(),
        Issuer:     "wordcount",
        Subject:    user.Name,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    tokenString, err := token.SignedString(jwtPrivateKey)
    if err != nil {
        log.Printf("JWT signing error: %s", err.Error())
        RenderJsonError(w, JsonError{ Status: http.StatusInternalServerError, Message: "Internal error" })
        return
    }

    RenderJson(w, http.StatusOK, JwtToken{ Token: tokenString })
}
