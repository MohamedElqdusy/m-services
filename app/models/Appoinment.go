package models

import (
	"time"
)

type Appoinment struct {
	Doctor    int       `json:"doctor"`
	TimePoint time.Time `json:"timePoint"`
	Patient   int       `json:patient`
	Notes     string    `json:"notes"`
}
