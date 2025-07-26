package models

import "time"

type Todo struct {
	ID        string
	Title     string
	Completed bool
	UserID    string
	CreatedAt time.Time
}
