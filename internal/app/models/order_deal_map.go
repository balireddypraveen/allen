package models

import (
	"github.com/google/uuid"
	"time"
)

type OrderDealMap struct {
	OrderID   uuid.UUID  `json:"order_id" gorm:"primaryKey"`
	DealId    uuid.UUID  `json:"deal_id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:current_timestamp;OnUpdate:SET current_timestamp;"`
}

func (o OrderDealMap) TableName() string {
	return "order_deal_map"
}
