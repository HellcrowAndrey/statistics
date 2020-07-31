package dto

import (
	"time"
)

type PurchaseDto struct {
	Id uint
	CreatedAt time.Time
	OrderId uint
	AccountId uint
	Products  []*ProductDto
}
