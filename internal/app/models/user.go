package models

import (
	"time"
)

type User struct {
	UserId    int64      `json:"user_id" gorm:"type:int;not null"`
	Name      string     `json:"name" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:current_timestamp;OnUpdate:SET current_timestamp;"`
}

func (o User) TableName() string {
	return "users"
}
