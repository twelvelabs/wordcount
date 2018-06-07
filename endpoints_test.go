package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_CreateTokenEndpoint_garbage(t *testing.T) {
    reqJson := `}$!.this."is._garbage`
    req, err := http.NewRequest("POST", "/token", strings.NewReader(reqJson))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    http.HandlerFunc(CreateTokenEndpoint).ServeHTTP(rr, req)

    // Check the status code is what we expect.
    assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_CreateTokenEndpoint_invalid_json(t *testing.T) {
    reqJson := `{ "lol": "not a user" }`
    req, err := http.NewRequest("POST", "/token", strings.NewReader(reqJson))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    http.HandlerFunc(CreateTokenEndpoint).ServeHTTP(rr, req)

    assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func Test_CreateTokenEndpoint_invalid_password(t *testing.T) {
    user := NewUserService().users[0]
    reqJson := fmt.Sprintf(`{ "name": "%s", "password": "nope" }`, user.Name)
    req, err := http.NewRequest("POST", "/token", strings.NewReader(reqJson))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    http.HandlerFunc(CreateTokenEndpoint).ServeHTTP(rr, req)

    assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func Test_CreateTokenEndpoint_valid_username_and_password(t *testing.T) {
    user := NewUserService().users[0]
    reqJson := fmt.Sprintf(`{ "name": "%s", "password": "%s" }`, user.Name, user.Password)
    req, err := http.NewRequest("POST", "/token", strings.NewReader(reqJson))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    http.HandlerFunc(CreateTokenEndpoint).ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Equal(t, "{\"token\":\"lolwat\"}\n", rr.Body.String())
}
