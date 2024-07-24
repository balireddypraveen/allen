package dto

import "github.com/google/uuid"

type GetDealsRequest struct {
	UserId  int64
	OrderId uuid.UUID
}

type CreateDealRequest struct {
	DealName    string `json:"deal_name" binding:"required"`
	MaxQuantity int    `json:"max_quantity" binding:"required"`
}

type UpdateDealRequest struct {
	DealId                 uuid.UUID `json:"deal_id" binding:"required"`
	MaxQuantity            int       `json:"max_quantity"`
	ExtendEndTimeByMinutes int       `json:"extend_end_time_by_min"`
}

type EndDealRequest struct {
	OrderId uuid.UUID `json:"order_id" binding:"required"`
}
