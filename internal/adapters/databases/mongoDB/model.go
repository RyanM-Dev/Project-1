package mongoDB

import "time"

type book struct {
	ID        string
	Title     string     `bson:"title,omitempty"`
	Author    string     `bson:"author,omitempty"`
	Published *time.Time `bson:"published,omitempty"`
	Genre     string     `bson:"genre,omitempty"`
}
