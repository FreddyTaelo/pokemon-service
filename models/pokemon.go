// models/pokemon.go
package models

type Pokemon struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Image  string   `json:"image"`
	Height int      `json:"height"`
	Weight int      `json:"weight"`
	Types  []string `json:"types"`
	Stats  []Stat   `json:"stats"`
}

type Stat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
