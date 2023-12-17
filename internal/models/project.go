package models

type Student struct {
	Name          string `json:"name"`
	Enrollment_No string `json:"Enrollment No"`
	Email         string `json:"email"`
	Reason        string `json:"Reason"`
}
