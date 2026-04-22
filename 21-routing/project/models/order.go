package models

import "time"

type Order struct {
	ID        string    `json:"id"`
	Products  []Product `json:"products"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Total     float64   `json:"total"`
}
