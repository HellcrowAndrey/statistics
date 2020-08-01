package dto

import (
	"time"
)

// PurchaseDto swagger:model
type PurchaseDto struct {
	Id        uint          `json:"id" example:"1"`
	CreatedAt time.Time     `json:"createAt" example:"2019-11-09T21:21:46+00:00"`
	OrderId   uint          `json:"orderId" example:"1"`
	AccountId uint          `json:"accountId omitempty" example:"1"`
	Products  []*ProductDto `json:"products"`
}
