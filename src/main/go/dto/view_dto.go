package dto

import "time"

type ViewDto struct {
	Id        uint
	CreatedAt time.Time
	Products  []*ProductDto
}
