package dto

import (
	"time"
)

type ViewDto struct {
	Id        uint          `json:"id"`
	CreatedAt time.Time     `json:"createAt"`
	AccountId int64         `json:"accountId,omitempty"`
	Products  []*ProductDto `json:"products"`
}
