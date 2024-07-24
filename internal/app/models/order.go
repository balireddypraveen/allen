package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OrderID   uuid.UUID  `json:"order_id" gorm:"primaryKey"`
	UserId    int        `json:"user_id" gorm:"type:int;not null"`
	DealId    uuid.UUID  `json:"deal_id" gorm:"type:int;not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:current_timestamp;OnUpdate:SET current_timestamp;"`
}

func (o Order) TableName() string {
	return "orders"
}
