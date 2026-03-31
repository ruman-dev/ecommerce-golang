package models

import "time"

type User struct {
	ID         string    `json:"id" db:"id"`
	Name       string    `db:"name" json:"name"`
	Age        int       `db:"age" json:"age"`
	Email      string    `db:"email" json:"email"`
	Password   string    `db:"password" json:"-"`
	Created_At time.Time `db:"created_at" json:"created_at"`
}
