package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte
}

type Response struct {
	Message string `json: "message"`
	Success bool   `json: "success"`
}