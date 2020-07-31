package dto

import (
	"time"
)

type PurchaseDto struct {
	Id        uint          `json:"id"`
	CreatedAt time.Time     `json:"createAt"`
	OrderId   uint          `json:"orderId"`
	AccountId uint          `json:"accountId,omitempty"`
	Products  []*ProductDto `json:"products"`
}
