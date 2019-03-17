package models

type Car struct {
	Code string    `json:"code"`
	Make  string    `json:"make"`
	Model string 	`json:"model"`
	Kw    int64       `json:"kw"`
	Year  int64       `json:"year"`
	Color string 	`json:"color"`
	Price int64 		`json:"price"`
}
