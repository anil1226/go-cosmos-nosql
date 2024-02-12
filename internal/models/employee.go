package models

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Branch Branch `json:"branch"`
}

type Branch struct {
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zip"`
}
