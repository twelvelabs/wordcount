package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
    "github.com/jdkato/prose/tokenize"
)


type JsonError struct {
    Status  int     `json:"status"`
    Message string  `json:"message"`
}

type JwtToken struct {
    Token   string  `json:"token"`
}

type WordcountResponse struct {
    Count   int             `json:"count"`
    Words   map[string]int  `json:"words"`
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

    found, err := NewUserService().AuthenticateCredentials(user.Name, user.Password)
    if err != nil {
        RenderJsonError(w, JsonError{ Status: http.StatusUnauthorized, Message: "Invalid credentials" })
        return
    }

    RenderJson(w, http.StatusOK, JwtToken{ Token: found.GenerateToken() })
}

func WordcountEndpoint(w http.ResponseWriter, r *http.Request) {
    extractor := request.AuthorizationHeaderExtractor
    token, err := request.ParseFromRequest(r, extractor, func(token *jwt.Token) (interface{}, error) {
        return jwtPublicKey, nil
    })
    if err != nil || !token.Valid {
        RenderJsonError(w, JsonError{ Status: http.StatusUnauthorized, Message: "Invalid token" })
        return
    }

    // neither `strings.ToLower` nor the tokenizer accept an io.Reader,
    // so we need to copy the request body over to a string :(
    buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
    body := buf.String()
    // rip the body string into an array of lowercase tokens...
    tokenizer := tokenize.NewWordBoundaryTokenizer()
    tokens := tokenizer.Tokenize(strings.ToLower(body))
    // count up how many times each token is present...
    words := make(map[string]int)
    for _, t := range tokens {
        words[t] += 1
    }
    // and wrap it up in a response object!
    wr := WordcountResponse{
        Count: len(tokens),
        Words: words,
    }
    RenderJson(w, http.StatusOK, wr)
}

