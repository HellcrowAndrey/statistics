package dto

import (
	"time"
)

type AccountDto struct {
	Id        uint           `json:"id"`
	CreatedAt time.Time      `json:"createAt"`
	UserId    uint           `json:"userId"`
	Purchases []*PurchaseDto `json:"purchases"`
	Logins    []*LoginDto    `json:"logins"`
	Views     []*ViewDto     `json:"views"`
}
