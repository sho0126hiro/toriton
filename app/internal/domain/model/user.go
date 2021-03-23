package model

import (
	"time"
)

type User struct {
	Id              string
	Name            string
	FirebaseId      string
	ProfileImageUrl string
	Biography       string
	Birthday        string
	Email           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
