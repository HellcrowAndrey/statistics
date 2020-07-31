package dto

import (
	"time"
)

type LoginDto struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	AccountId int64     `json:"accountId"`
	Ip        string    `json:"ip"`
	Device    string    `json:"device"`
	Location  string    `json:"location"`
}
