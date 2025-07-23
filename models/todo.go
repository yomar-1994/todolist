package models

import "time"

type Todo struct {
	ID        string
	title     string
	completed bool
	UserID    string
	createdAt time.Time
}
