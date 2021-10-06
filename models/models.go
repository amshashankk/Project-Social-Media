package models

type User struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password []byte
}

type Response struct {
	Message string `json: "message"`
	Success bool   `json: "success"`
}
