package dto

import (
	"time"
)

type LoginDto struct {
	Id        uint
	CreatedAt time.Time
	AccountId int64
	Ip        string
	Device    string
	Location  string
}
