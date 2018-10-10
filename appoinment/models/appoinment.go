package models

import "time"

type Appoinment struct {
	DoctorId  int       `json:"doctor"`
	TimePoint time.Time `json:"timepoint"`
	PatientId int       `json:"patient"`
	Notes     string    `json:"notes"`
}
