package models

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Reason string `json:"reason"`
}
