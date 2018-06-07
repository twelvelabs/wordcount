package main

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "log"
)

type User struct {
    Name string
    Password string
    Token string
}

type UserService struct {
    users []User
}

// Low-fi user storage and authentication (I wanted to keep this simple).
// In a real world system, we'd probably be using a DB (w/ hashed passwords)
// or some other centralized service.
func (s UserService) AuthenticateCredentials(name string, password string) (user User, err error) {
    for _, u := range s.users {
      if u.Name == name && u.Password == password {
          return u, nil
      }
      if u.Name == name {
          return user, errors.New("invalid password")
      }
    }
    return user, errors.New("invalid username")
}

func NewUserService() UserService {
    content, err := ioutil.ReadFile("/home/wordcount/users.json")
    if err != nil {
      log.Fatal(err)
    }

    var users []User
    json.Unmarshal(content, &users)

    return UserService{ users: users }
}
