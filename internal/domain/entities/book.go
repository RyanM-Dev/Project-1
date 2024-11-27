package entities

import "time"

type Book struct {
	ID        string
	Title     string
	Author    string
	Published *time.Time
	Genre     string
}
