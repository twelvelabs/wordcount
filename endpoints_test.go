package main

import (
    "encoding/json"
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
}

func Test_WordcountEndpoint_empty_body(t *testing.T) {
    reqText := ""
    req, err := http.NewRequest("POST", "/wordcount", strings.NewReader(reqText))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    http.HandlerFunc(WordcountEndpoint).ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var wcRes WordcountResponse
    err = json.NewDecoder(rr.Body).Decode(&wcRes)
    assert.NoError(t, err)

    assert.Equal(t, 0, wcRes.Count)
    assert.Equal(t, 0, len(wcRes.Words))
}

func Test_WordcountEndpoint_simple_text(t *testing.T) {
    reqText := `I don't know why you say "Goodbye", I say "Hello, hello, hello".`
    req, err := http.NewRequest("POST", "/wordcount", strings.NewReader(reqText))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    http.HandlerFunc(WordcountEndpoint).ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var wcRes WordcountResponse
    err = json.NewDecoder(rr.Body).Decode(&wcRes)
    assert.NoError(t, err)

    assert.Equal(t, 12, wcRes.Count)
    assert.Equal(t, 2, wcRes.Words["i"])
    assert.Equal(t, 1, wcRes.Words["don't"])
    assert.Equal(t, 1, wcRes.Words["know"])
    assert.Equal(t, 1, wcRes.Words["you"])
    assert.Equal(t, 2, wcRes.Words["say"])
    assert.Equal(t, 1, wcRes.Words["goodbye"])
    assert.Equal(t, 3, wcRes.Words["hello"])
}
