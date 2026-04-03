package models

import "time"

type Product struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Price       float64   `db:"price" json:"price"`
	Quantity    int       `db:"quantity" json:"quantity"`
	Description string    `db:"description" json:"description"`
	Created_At  time.Time `db:"created_at" json:"created_at"`
}
