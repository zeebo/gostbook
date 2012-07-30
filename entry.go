package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Entry struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Timestamp time.Time
	Name      string
	Message   string
}

func NewEntry() *Entry {
	return &Entry{
		Timestamp: time.Now(),
	}
}
