package models

import "time"

type Patient struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Id        int       `json:"id"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
