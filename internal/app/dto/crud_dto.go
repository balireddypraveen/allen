package dto

import "github.com/google/uuid"

type GetOrdersRequest struct {
	UserId  int64
	OrderId uuid.UUID
}

type CreateOrderRequest struct {
	UserId int       `json:"user_id" binding:"required"`
	DealId uuid.UUID `json:"deal_id" binding:"required"`
}

type CancelOrderRequest struct {
	OrderId uuid.UUID `json:"order_id" binding:"required"`
}
