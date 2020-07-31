package dto

import (
	"time"
)

type LoginDto struct {
	Id        uint
	CreatedAt time.Time
	Ip        string
	Device    string
	Location  string
}
