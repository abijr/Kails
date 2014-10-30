package models

import "time"

type Flashcard struct {
	Definition string
	Lesson     int
	Bucket     int
	LastReview time.Time
}
