package models

import (
	"github.com/google/uuid"
	"time"
)

type Deal struct {
	DealId      uuid.UUID  `json:"deal_id" gorm:"primaryKey"`
	DealName    string     `json:"deal_name" gorm:"primaryKey"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	MaxQuantity int        `json:"max_quantity"`
	Enabled     bool       `json:"enabled"`
	CreatedAt   *time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"default:current_timestamp;OnUpdate:SET current_timestamp;"`
}

func (o Deal) TableName() string {
	return "deals"
}
