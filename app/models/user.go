package models

type UserType int

const (
	PatientType UserType = iota
	DoctorType
)

type User struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	UserType  UserType `json:"userType"`
	Password  string   `json:"password"`
}
