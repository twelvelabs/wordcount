package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_NewUserService(t *testing.T) {
    service := NewUserService()
    user := service.users[0]
    assert.Equal(t, "bob", user.Name)
}

func Test_UserService_AuthenticateCredentials_invalid_name(t *testing.T) {
    service := NewUserService()

    found, err := service.AuthenticateCredentials("lol", "wat")
    assert.Error(t, err)
    assert.IsType(t, User{}, found)
}

func Test_UserService_AuthenticateCredentials_invalid_pass(t *testing.T) {
    service := NewUserService()

    found, err := service.AuthenticateCredentials("bob", "nope")
    assert.Error(t, err)
    assert.IsType(t, User{}, found)
}

func Test_UserService_AuthenticateCredentials_valid_user_and_pass(t *testing.T) {
    service := NewUserService()
    user := service.users[0]
    assert.Equal(t, "bob", user.Name)

    found, err := service.AuthenticateCredentials("bob", user.Password)
    assert.NoError(t, err)
    assert.Equal(t, user, found)
}
