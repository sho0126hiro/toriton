package dao

import (
	"time"
)

type User struct {
	Id              int64     `json:"id" gorm:"primary_key;unique"`
	FirebaseId      string    `json:"firebase_id"`
	Username        string    `json:"username" gorm:"unique"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Biography       string    `json:"biography"`
	ProfileImageUrl string    `json:"profile_image_url"`
	Birthday        string    `json:"birthday"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
