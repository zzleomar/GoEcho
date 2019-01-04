package models

type User struct {
	IdUser string `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Email  string `json:"email" xml:"email"`
}
