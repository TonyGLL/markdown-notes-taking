package db

import "time"

type Note struct {
	ID        int       `json:"id"`
	HTML      string    `json:"html"`
	MK        string    `json:"mk"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"deleted"`
}
