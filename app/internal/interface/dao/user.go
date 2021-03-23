package dao

import (
	"time"
)

type User struct {
	Id              string    `json:"id" gorm:"primary_key;unique"`
	Name            string    `json:"name"`
	FirebaseId      string    `json:"firebase_id"`
	ProfileImageUrl string    `json:"profile_image_url"`
	Biography       string    `json:"biography"`
	Birthday        string    `json:"birthday"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
