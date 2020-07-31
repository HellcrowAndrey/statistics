package dto

import (
	"time"
)

type ViewDto struct {
	Id        uint
	CreatedAt time.Time
	AccountId int64
	Products  []*ProductDto
}
