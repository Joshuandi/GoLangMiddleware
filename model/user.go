package model

import "time"

var users = []*User{}

type User struct {
	Id         string    `json: "id"`
	Username   string    `json: "username"`
	Email      string    `json: "email"`
	Password   string    `json: "password"`
	Age        int       `json: "age"`
	Division   string    `json: "division"`
	Created_at time.Time `json : "created_at"`
	Updated_at time.Time `json : "updated_at"`
}
